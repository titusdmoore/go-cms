// Package templates handles the file IO and template/html parsing of views.
package templates

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Templates struct {
	templates *template.Template
}

type TemplatePageData struct {
	Data interface{}

	actions []interface{}
	filters []interface{}
}

func (pd *TemplatePageData) doAction(action string, opts ...interface{}) string {
	log.Println("This is running")
	return "Testing Stuff"
}

// Here is what I am thinking
type PageDataWithActions interface {
	doAction(action string, opts ...interface{})
	applyFilter(filter string, opts ...interface{})
}

func (t *Templates) Render(writer io.Writer, name string, data TemplatePageData) error {
	// FuncMap Needed?
	return t.templates.ExecuteTemplate(writer, name, data)
}

func InitializeTemplates() *Templates {
	return &Templates{
		templates: template.Must(parseTemplateDirs("./internal/templates/views")),
	}
}

// parseTemplateDirs walks the passed path to pull parse all templates in nested heirarchy.
func parseTemplateDirs(path string) (*template.Template, error) {
	templ := template.New("")

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseGlob(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ, nil
}
