package main

import (
	"services/internal/config"
	"services/internal/database/postgresql"
)

func main() {
	environment, err := config.Load[config.CoreEnvironment]()
	if err != nil {
		panic(err)
	}

	postgresql.LoadPostgresUri(environment.PostgresUri)
}
