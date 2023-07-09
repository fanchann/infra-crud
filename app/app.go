package app

import (
	"log"

	"github.com/zakirkun/infra-go/pkg/database"
	"github.com/zakirkun/infra-go/pkg/server"

	"infra-book-crud/models"
)

type App interface {
	Boot()
}

type appCtx struct {
	db     database.DBModel
	server server.ServerContext
}

func NewApp(db database.DBModel, srv server.ServerContext) App {
	return appCtx{db: db, server: srv}
}

func (a appCtx) Boot() {
	// open connection
	_, err := a.db.OpenDB()

	if err != nil {
		log.Fatalf("Failed connect database : %v", err)
	}

	database.DB = &a.db

	// migration
	database.NewMigration(&models.Books{})

	// running application
	a.server.Run()
}
