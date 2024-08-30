package dbrepo

import (
	"database/sql"
	"PSInventory/internal/database"
)

type postgresDBRepo struct {
	DB  *sql.DB
}

func NewDBRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB:  conn,
	}
}