package main

import (
	// "io"
	"fmt"
	"net/http"

	"github.com/titusdmoore/goCms/internal/app"
	"github.com/titusdmoore/goCms/internal/components"
	"github.com/titusdmoore/goCms/internal/templates"
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
		host := "https://" + r.Host
		data := struct {
			PageTitle string
			Host      string
		}{PageTitle: "Welcome to GO CMS", Host: host}
		pageData := templates.NewTemplatePageData()
		pageData.Data = data

		push, ok := w.(http.Pusher)
		fmt.Println(ok)
		if ok {
			pushErr := push.Push(host+"/static/css/main.css", nil)
			pushErr = push.Push(host+"/static/js/htmx.min.js", nil)

			if pushErr != nil {
				fmt.Println(pushErr)
			}
		}

		err := components.Index(pageData, app).Render(r.Context(), w)
		if err != nil {
			panic(err)
		}
	})

	app.Router.RegisterRoute(app.Config.Router.AdminPath+"/pages", func(w http.ResponseWriter, r *http.Request) {
		host := "https://" + r.Host

		fmt.Println(host)

		pageData := PageData{
			PageTitle:      "View All Pages",
			Host:           host,
			PageName:       "Pages",
			SinglePageName: "Page",
			InternalAction: Action{
				name: "Testing",
			},
		}
		templateData := templates.NewTemplatePageData()
		templateData.Data = pageData

		app.Templates.Render(w, "table", templateData)
	})
	app.Router.RegisterRoute(app.Config.Router.AdminPath+"/pages/new", func(w http.ResponseWriter, r *http.Request) {
		app.Templates.Render(w, "new", templates.DefaultTemplateData())
	})

	app.EventManager.AddAction("header", 10, func(args ...any) {
		fmt.Println("This is my wonderful action")
	})

	app.EventManager.AddFilter("header", 10, func(args ...any) (any, error) {
		value := args[0].(string)
		return value + " - Filtered", nil
	})

	app.Router.Serve(app.Config)
}
