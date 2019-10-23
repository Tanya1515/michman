package database

import (
	proto "gitlab.at.ispras.ru/openstack_bigdata_tools/spark-openstack/src/protobuf"
)

type Database interface {
	WriteCluster(cluster *proto.Cluster) error
	ReadCluster(name string) (*proto.Cluster, error)
	ListClusters() ([]proto.Cluster, error)
	DeleteCluster(name string) error

	ListProjects() ([]proto.Project, error)
	ReadProject(name string) (*proto.Project, error)
	WriteProject(project *proto.Project) error
	ReadProjectClusters(projectID string) ([]proto.Cluster, error)
	UpdateProject(*proto.Project) error
	DeleteProject(name string) error

	ReadTemplate(projectID, id string) (*proto.Template, error)
	ReadTemplateByName(templateName string) (*proto.Template, error)
	WriteTemplate(template *proto.Template) error
	DeleteTemplate(id string) error
	ListTemplates(projectID string) ([]proto.Template, error)
}
