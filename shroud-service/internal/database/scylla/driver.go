package scylla

import "github.com/gocql/gocql"

var cluster *gocql.ClusterConfig

func LoadCluster(scyllaUri string) {
	cluster = gocql.NewCluster(scyllaUri)
}

func NewSession() (*gocql.Session, error) {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}
