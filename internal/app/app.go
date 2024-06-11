package app

import (
	"sync"

	"github.com/titusdmoore/goCms/internal/config"
	"github.com/titusdmoore/goCms/internal/db"
	"github.com/titusdmoore/goCms/internal/events"
	"github.com/titusdmoore/goCms/internal/router"
	"github.com/titusdmoore/goCms/internal/templates"
)

var lock = &sync.Mutex{}

type application struct {
	Database     db.DB
	Config       config.Config
	Router       router.Router
	Templates    *templates.Templates
	EventManager events.EventManager
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

	app.Templates = templates.InitializeTemplates()
	app.EventManager = events.InitializeEventManager()

	return app
}
