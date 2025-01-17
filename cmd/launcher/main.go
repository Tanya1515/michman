package main

import (
	"flag"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/ispras/michman/internal/database"
	"github.com/ispras/michman/internal/protobuf"
	"github.com/ispras/michman/internal/utils"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

const (
	launcherDefaultPort = "5000"
)

type ansibleLaunch interface {
	Run(c *protobuf.Cluster, osCreds *utils.OsCredentials, dockRegCreds *utils.DockerCredentials, osConfig *utils.Config, action string) string
}

type ansibleService struct {
	logger            *log.Logger
	ansibleRunner     ansibleLaunch
	vaultCommunicator utils.SecretStorage
	config            utils.Config
}

func (aS *ansibleService) Init(logger *log.Logger, ansibleLaunch AnsibleLauncher,
	vaultCommunicator utils.SecretStorage, config *utils.Config) {
	aS.logger = logger
	aS.ansibleRunner = ansibleLaunch
	aS.vaultCommunicator = vaultCommunicator
	aS.config = *config
}

func makeOsCreds(keyName string, vaultClient *vaultapi.Client, version string) *utils.OsCredentials {
	secretValues, err := vaultClient.Logical().Read(keyName)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var osCreds utils.OsCredentials
	switch version {
	case utils.OsUssuriVersion:
		osCreds.OsAuthUrl = secretValues.Data[utils.OsAuthUrl].(string)
		osCreds.OsProjectName = secretValues.Data[utils.OsProjectName].(string)
		osCreds.OsProjectID = secretValues.Data[utils.OsProjectID].(string)
		osCreds.OsInterface = secretValues.Data[utils.OsInterface].(string)
		osCreds.OsPassword = secretValues.Data[utils.OsPassword].(string)
		osCreds.OsRegionName = secretValues.Data[utils.OsRegionName].(string)
		osCreds.OsUserName = secretValues.Data[utils.OsUsername].(string)
		osCreds.OsUserDomainName = secretValues.Data[utils.OsUserDomainName].(string)
		osCreds.OsProjectDomainID = secretValues.Data[utils.OsProjectDomainID].(string)
		osCreds.OsIdentityApiVersion = secretValues.Data[utils.OsIdentityApiVersion].(string)
	case utils.OsSteinVersion:
		osCreds.OsAuthUrl = secretValues.Data[utils.OsAuthUrl].(string)
		osCreds.OsPassword = secretValues.Data[utils.OsPassword].(string)
		osCreds.OsProjectName = secretValues.Data[utils.OsProjectName].(string)
		osCreds.OsRegionName = secretValues.Data[utils.OsRegionName].(string)
		osCreds.OsUserName = secretValues.Data[utils.OsUsername].(string)
		osCreds.OsComputeApiVersion = secretValues.Data[utils.OsComputeApiVersion].(string)
		osCreds.OsNovaVersion = secretValues.Data[utils.OsNovaVersion].(string)
		osCreds.OsAuthType = secretValues.Data[utils.OsAuthType].(string)
		osCreds.OsCloudname = secretValues.Data[utils.OsCloudname].(string)
		osCreds.OsIdentityApiVersion = secretValues.Data[utils.OsIdentityApiVersion].(string)
		osCreds.OsImageApiVersion = secretValues.Data[utils.OsImageApiVersion].(string)
		osCreds.OsNoCache = secretValues.Data[utils.OsNoCache].(string)
		osCreds.OsProjectDomainName = secretValues.Data[utils.OsProjectDomainName].(string)
		osCreds.OsUserDomainName = secretValues.Data[utils.OsUserDomainName].(string)
		osCreds.OsVolumeApiVersion = secretValues.Data[utils.OsVolumeApiVersion].(string)
		osCreds.OsPythonwarnings = secretValues.Data[utils.OsPythonwarnings].(string)
		osCreds.OsNoProxy = secretValues.Data[utils.OsNoProxy].(string)
	case utils.OsLibertyVersion:
		osCreds.OsAuthUrl = secretValues.Data[utils.OsAuthUrl].(string)
		osCreds.OsPassword = secretValues.Data[utils.OsPassword].(string)
		osCreds.OsProjectName = secretValues.Data[utils.OsProjectName].(string)
		osCreds.OsRegionName = secretValues.Data[utils.OsRegionName].(string)
		osCreds.OsTenantId = secretValues.Data[utils.OsTenantId].(string)
		osCreds.OsTenantName = secretValues.Data[utils.OsTenantName].(string)
		osCreds.OsUserName = secretValues.Data[utils.OsUsername].(string)
		if uname, ok := secretValues.Data[utils.OsSwiftUsername]; ok {
			osCreds.OsSwiftUserName = uname.(string)
		} else {
			osCreds.OsSwiftUserName = ""
		}
		if pass, ok := secretValues.Data[utils.OsSwiftPassword]; ok {
			osCreds.OsSwiftUserName = pass.(string)
		} else {
			osCreds.OsSwiftPassword = ""
		}
	default: //liberty as default version
		osCreds.OsAuthUrl = secretValues.Data[utils.OsAuthUrl].(string)
		osCreds.OsPassword = secretValues.Data[utils.OsPassword].(string)
		osCreds.OsProjectName = secretValues.Data[utils.OsProjectName].(string)
		osCreds.OsRegionName = secretValues.Data[utils.OsRegionName].(string)
		osCreds.OsTenantId = secretValues.Data[utils.OsTenantId].(string)
		osCreds.OsTenantName = secretValues.Data[utils.OsTenantName].(string)
		osCreds.OsUserName = secretValues.Data[utils.OsUsername].(string)
		if uname, ok := secretValues.Data[utils.OsSwiftUsername]; ok {
			osCreds.OsSwiftUserName = uname.(string)
		} else {
			osCreds.OsSwiftUserName = ""
		}
		if pass, ok := secretValues.Data[utils.OsSwiftPassword]; ok {
			osCreds.OsSwiftUserName = pass.(string)
		} else {
			osCreds.OsSwiftPassword = ""
		}
	}

	return &osCreds
}

func makeDockerCreds(keyName string, vaultClient *vaultapi.Client) *utils.DockerCredentials {
	secrets, err := vaultClient.Logical().Read(keyName)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var res utils.DockerCredentials
	res.Url = secrets.Data[utils.DockerLoginUlr].(string)
	res.User = secrets.Data[utils.DockerLoginUser].(string)
	res.Password = secrets.Data[utils.DockerLoginPassword].(string)
	return &res
}

func checkSshKey(keyName string, vaultClient *vaultapi.Client) error {

	sshPath := filepath.Join(utils.SshKeyPath)
	if _, err := os.Stat(sshPath); os.IsNotExist(err) {
		secretValues, err := vaultClient.Logical().Read(keyName)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		sshKey := secretValues.Data[utils.VaultSshKey].(string)
		f, err := os.Create(sshPath)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		err = os.Chmod(sshPath, 0777)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = f.WriteString(sshKey)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		err = f.Close()
		if err != nil {
			log.Fatalln(err)
			return err
		}
		err = os.Chmod(sshPath, 0400)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func (aS *ansibleService) Delete(in *protobuf.Cluster, stream protobuf.AnsibleRunner_DeleteServer) error {
	aS.logger.Print("Getting delete cluster request...")
	aS.logger.Print("Cluster info:")
	in.PrintClusterData(aS.logger)

	aS.logger.Print("Getting vault secrets...")

	vaultClient, vaultCfg := aS.vaultCommunicator.ConnectVault()
	if vaultClient == nil {
		log.Fatalln("Error: can't connect to vault secrets storage")
		return nil
	}

	keyName := vaultCfg.OsKey

	osCreds := makeOsCreds(keyName, vaultClient, aS.config.OsVersion)
	if osCreds == nil {
		return nil
	}

	err := checkSshKey(vaultCfg.SshKey, vaultClient)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	var dockRegCreds *utils.DockerCredentials
	if aS.config.SelfignedRegistry || aS.config.GitlabRegistry {
		dockRegCreds = makeDockerCreds(vaultCfg.RegistryKey, vaultClient)
	}

	ansibleStatus := aS.ansibleRunner.Run(in, osCreds, dockRegCreds, &aS.config, utils.ActionDelete)

	if err := stream.Send(&protobuf.TaskStatus{Status: ansibleStatus}); err != nil {
		return err
	}

	return nil
}

func (aS *ansibleService) Update(in *protobuf.Cluster, stream protobuf.AnsibleRunner_UpdateServer) error {
	aS.logger.Print("Getting update cluster request...")
	aS.logger.Print("Cluster info:")
	in.PrintClusterData(aS.logger)

	aS.logger.Print("Getting vault secrets...")

	vaultClient, vaultCfg := aS.vaultCommunicator.ConnectVault()
	if vaultClient == nil {
		log.Fatalln("Error: can't connect to vault secrets storage")
		return nil
	}

	osCreds := makeOsCreds(vaultCfg.OsKey, vaultClient, aS.config.OsVersion)
	if osCreds == nil {
		return nil
	}

	err := checkSshKey(vaultCfg.SshKey, vaultClient)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	var dockRegCreds *utils.DockerCredentials
	if aS.config.SelfignedRegistry || aS.config.GitlabRegistry {
		dockRegCreds = makeDockerCreds(vaultCfg.RegistryKey, vaultClient)
	}

	ansibleStatus := aS.ansibleRunner.Run(in, osCreds, dockRegCreds, &aS.config, utils.ActionUpdate)

	if err := stream.Send(&protobuf.TaskStatus{Status: ansibleStatus}); err != nil {
		return err
	}
	return nil
}

func (aS *ansibleService) Create(in *protobuf.Cluster, stream protobuf.AnsibleRunner_CreateServer) error {
	aS.logger.Print("Getting create cluster request...")
	aS.logger.Print("Cluster info:")
	in.PrintClusterData(aS.logger)

	aS.logger.Print("Getting vault secrets...")
	vaultClient, vaultCfg := aS.vaultCommunicator.ConnectVault()
	if vaultClient == nil {
		log.Fatalln("Error: can't connect to vault secrets storage")
		return nil
	}

	osCreds := makeOsCreds(vaultCfg.OsKey, vaultClient, aS.config.OsVersion)
	if osCreds == nil {
		return nil
	}

	err := checkSshKey(vaultCfg.SshKey, vaultClient)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var dockRegCreds *utils.DockerCredentials
	if aS.config.SelfignedRegistry || aS.config.GitlabRegistry {
		dockRegCreds = makeDockerCreds(vaultCfg.RegistryKey, vaultClient)
	}

	ansibleStatus := aS.ansibleRunner.Run(in, osCreds, dockRegCreds, &aS.config, utils.ActionCreate)

	if err := stream.Send(&protobuf.TaskStatus{Status: ansibleStatus}); err != nil {
		return err
	}

	return nil
}

func (aS *ansibleService) GetMasterIP(in *protobuf.Cluster, stream protobuf.AnsibleRunner_GetMasterIPServer) error {
	return nil
}

func main() {
	//set flags for config path and ansible service adress
	configPath := flag.String("config", utils.ConfigPath, "Path to the config.yaml file")
	launcherPort := flag.String("port", launcherDefaultPort, "Launcher service default port")
	flag.Parse()

	//set config file path
	utils.SetConfigPath(*configPath)
	config := utils.Config{}
	if err := config.MakeCfg(); err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile(config.LogsFilePath+"/launcher.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	ansibleServiceLogger := log.New(mw, "LAUNCHER: ", log.Ldate|log.Ltime)
	vaultCommunicator := utils.VaultCommunicator{}
	err = vaultCommunicator.Init()
	if err != nil {
		panic(err)
	}
	db, err := database.NewCouchBase(&vaultCommunicator)
	if err != nil {
		panic("Can't create database connection. Exit...")
	}
	ansibleLaunch := AnsibleLauncher{couchbaseCommunicator: db}
	lis, err := net.Listen("tcp", ":"+*launcherPort)
	if err != nil {
		ansibleServiceLogger.Fatalf("failed to listen: %v", err)
	}

	gas := grpc.NewServer()
	aService := ansibleService{}
	aService.Init(ansibleServiceLogger, ansibleLaunch, &vaultCommunicator, &config)
	protobuf.RegisterAnsibleRunnerServer(gas, &aService)

	ansibleServiceLogger.Print("Ansible runner start work...\n")
	if err := gas.Serve(lis); err != nil {
		ansibleServiceLogger.Fatalf("failed to serve: %v", err)
	}
}
