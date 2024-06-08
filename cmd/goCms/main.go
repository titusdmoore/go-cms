package main

import (
	// "io"
	"net/http"
	"strings"

	"github.com/titusdmoore/goCms/internal/app"
)

type Action struct {
	name string
}

func (a *Action) Execute() string {
	return a.name
}

type PageData struct {
	PageTitle      string
	Host           string
	PageName       string
	SinglePageName string
	InternalAction Action
}

func (p *PageData) TestingAction() string {
	return "This is a returned string"
}

func main() {
	app := app.InitializeProject()

	// fmt.Printf("\n%v\n", app.Config)
	//
	//	if err := app.Database.PurgeDatabase(); err != nil {
	//		panic(err)
	//	}
	//
	// body, err := os.ReadFile("./sql/init.sql")
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if _, err := app.Database.Db.Exec(string(body)); err != nil {
	//		panic(err)
	//	}

	// app.Router.RegisterRoute("/", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "Hello, World!\n")
	// })

	app.Router.RegisterRoute(app.Config.Router.AdminPath, func(w http.ResponseWriter, r *http.Request) {
		host := strings.ToLower(strings.Split(r.Proto, "/")[0]) + "://" + r.Host

		app.Templates.Render(w, "index", struct {
			PageTitle string
			Host      string
		}{PageTitle: "Welcome to GO CMS", Host: host})
	})

	app.Router.RegisterRoute(app.Config.Router.AdminPath+"/pages", func(w http.ResponseWriter, r *http.Request) {
		host := strings.ToLower(strings.Split(r.Proto, "/")[0]) + "://" + r.Host

		pageData := PageData{
			PageTitle:      "View All Pages",
			Host:           host,
			PageName:       "Pages",
			SinglePageName: "Page",
			InternalAction: Action{
				name: "Testing",
			},
		}

		app.Templates.Render(w, "table", pageData)
	})
	app.Router.RegisterRoute(app.Config.Router.AdminPath+"/pages/new", func(w http.ResponseWriter, r *http.Request) {
		app.Templates.Render(w, "new", nil)
	})

	app.Router.Serve(app.Config)
}
