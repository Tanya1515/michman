package handlers

import (
	"encoding/json"
	"errors"
	protobuf "github.com/ispras/michman/internal/protobuf"
	"log"
	"strconv"
)

func deleteSpaces(valStr string) string {
	resStr := ""
	for _, ch := range valStr {
		if ch != ' ' {
			resStr += string(ch)
		}
	}
	return resStr
}

func checkValuesAllowed(val string, posVal []string) bool {
	val = deleteSpaces(val)
	for _, pv := range posVal {
		if val == pv {
			return true
		}
	}
	return false
}

func ValidateService(hS HttpServer, service *protobuf.Service) (bool, error) {
	hS.Logger.Print("Validating service type and config params...")

	if service.Type == "" {
		log.Print("ERROR: service type can't be nil.")
		return false, errors.New("ERROR: service type can't be nil.")
	}

	sTypes, err := hS.Db.ListServicesTypes()
	if err != nil {
		log.Print(err)
		return false, err
	}

	//check that service type is supported
	stOk := false
	var stIdx int
	for i, st := range sTypes {
		if st.Type == service.Type {
			stOk = true
			stIdx = i
			break
		}
	}

	if !stOk {
		log.Print("ERROR: service type ", service.Type, " is not supported.")
		return false, errors.New("ERROR: service type " + service.Type + " is not supported.")
	}

	//check service version
	if service.Version == "" && sTypes[stIdx].DefaultVersion != "" {
		service.Version = sTypes[stIdx].DefaultVersion
	} else if service.Version == "" && sTypes[stIdx].DefaultVersion == "" {
		log.Print("ERROR: service version and default version for service type " + service.Type + " are nil.")
		return false, errors.New("ERROR: service version and default version for service type " + service.Type + " are nil.")
	}

	//get idx of service version
	var svIdx int
	svOk := false
	for i, sv := range sTypes[stIdx].Versions {
		if sv.Version == service.Version {
			svIdx = i
			svOk = true
			break
		}
	}

	if !svOk {
		log.Print("ERROR: service version ", service.Version, " is not supported.")
		return false, errors.New("ERROR: service version " + service.Version + " is not supported.")
	}

	//validate configs
	for k, v := range service.Config {
		flagPN := false
		for _, sc := range sTypes[stIdx].Versions[svIdx].Configs {
			if k == sc.ParameterName {
				flagPN = true

				//check type
				if !sc.IsList {
					switch sc.Type {
					case "int":
						if _, err := strconv.ParseInt(v, 10, 32); err != nil {
							log.Print(err)
							return false, err
						}
					case "float":
						if _, err := strconv.ParseFloat(v, 64); err != nil {
							log.Print(err)
							return false, err
						}
					case "bool":
						if _, err := strconv.ParseBool(v); err != nil {
							log.Print(err)
							return false, err
						}
					}
				} else {
					switch sc.Type {
					case "int":
						var valList []int64
						if err := json.Unmarshal([]byte(v), &valList); err != nil {
							log.Print(err)
							return false, err
						}
					case "float":
						var valList []float64
						if err := json.Unmarshal([]byte(v), &valList); err != nil {
							log.Print(err)
							return false, err
						}
					case "bool":
						var valList []bool
						if err := json.Unmarshal([]byte(v), &valList); err != nil {
							log.Print(err)
							return false, err
						}
					case "string":
						var valList []string
						if err := json.Unmarshal([]byte(v), &valList); err != nil {
							log.Print(err)
							return false, err
						}
					}
				}

				//check for possible values
				if sc.PossibleValues != nil {
					flagPV := checkValuesAllowed(v, sc.PossibleValues)
					if !flagPV {
						log.Print("ERROR: service config param value ", v, " is not supported.")
						return false, errors.New("ERROR: service version " + v + " is not supported.")
					}
				}

				break
			}
		}
		if !flagPN {
			log.Print("ERROR: service config param name ", k, " is not supported.")
			return false, errors.New("ERROR: service config param name " + k + " is not supported.")
		}
	}

	return true, nil
}
