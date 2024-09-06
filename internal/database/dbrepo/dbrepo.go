package dbrepo

import (
	repository "PSInventory/internal/database"
	"database/sql"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewDBRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}
