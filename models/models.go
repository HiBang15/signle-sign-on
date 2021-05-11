package models

import (
	"database/sql"
	"github.com/HiBang15/signle-sign-on.git/database/sqlc"
)

type Connector struct {
	db *sql.DB
	*db.Queries
}
