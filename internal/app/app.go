package app

import (
	"sync"

	"github.com/titusdmoore/goCms/internal/config"
	"github.com/titusdmoore/goCms/internal/db"
	"github.com/titusdmoore/goCms/internal/router"
)

var lock = &sync.Mutex{}

type application struct {
	Database db.DB
	Config   config.Config
	Router   router.Router
}

var applicationInstance *application

func GetApp() *application {
	if applicationInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if applicationInstance == nil {
			applicationInstance = &application{}
		}
	}

	return applicationInstance
}

func InitializeProject() *application {
	app := GetApp()

	config, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}
	app.Config = *config

	app.Database = *db.InitializeDatabaseConnection(app.Config)
	router, err := router.InitializeRouter()
	if err != nil {
		panic(err)
	}
	app.Router = router

	return app
}
