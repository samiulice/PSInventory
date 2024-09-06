package main

import (
	"PSInventory/backend"
	repository "PSInventory/internal/database"
	"PSInventory/internal/database/dbrepo"
	"PSInventory/internal/driver"
	"context"
	"database/sql"
)

// App struct
type App struct {
	ctx    context.Context
	DB     repository.DatabaseRepo
	DBpool *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	//Connection to database
	dbConn, err := driver.ConnectDB("host=localhost port=5432 dbname=psinventory user=postgres password=psi@2024 sslmode=disable")
	if err != nil {
		return
	}
	a.DBpool = dbConn
	a.DB = dbrepo.NewDBRepo(dbConn)

	//run backend
	err = backend.RunServer()
	if err != nil {
		return
	}

}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	a.DBpool.Close() //Close db connectivty
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
// func (a *App) Greet() []*models.Product {
// 	// brands, err := a.DB.
// 	var b []*models.Product
// 	return b
// }
