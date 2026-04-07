package main

import (
	"services/internal/config"
	"services/internal/database/scylla"
)

func main() {
	environment, err := config.Load[config.ScribeEnvironment]()
	if err != nil {
		panic(err)
	}

	scylla.LoadCluster(environment.ScyllaUri)
}
