package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var uri string

func LoadPostgresUri(envPostgresUri string) {
	uri = envPostgresUri
}

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
