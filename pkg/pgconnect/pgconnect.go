package pgconnect

import (
	"database/sql"
	"fmt"
)

type SslMode string

const (
	SslMode_Prefer  = "prefer"
	SslMode_Require = "require"
	SslMode_Disable = "disable"
)

type Dsn struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     int
	SslMode  SslMode
}

func (d Dsn) String() string {
	conStr := "postgresql://%s:%s@%s:%d/%s?sslmode=%s"
	return fmt.Sprintf(conStr, d.User, d.Password, d.Host, d.Port, d.DbName, d.SslMode)
}

func Connect(dsn Dsn) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}
