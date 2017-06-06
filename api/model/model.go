package model

import (
	"github.com/jmoiron/sqlx"
)

// DBPool database connection pool
type DBPool struct {
	Master *sqlx.DB
}
