package main

import (
	"fmt"
	// "os"

	"html/template"
	"io"
	"net/http"

	"github.com/titusdmoore/goCms/internal/app"
)

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

	app.Router.RegisterRoute("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	})

	app.Router.RegisterRoute(app.Config.Router.AdminPath, func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./internal/router/src/views/index.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w, "testing")
		if err != nil {
			panic(err)
		}

		fmt.Println("Gotten here")
	})

	app.Router.Serve(app.Config)
}
