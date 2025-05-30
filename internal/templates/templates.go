// Package templates handles the file IO and template/html parsing of views.
package templates

import (
	"fmt"
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

type Action struct {
	Action   func(args ...any)
	Priority int32
}

type TemplatePageData struct {
	Data any

	actions map[string][]Action
	filters map[string][]Action
}

func (pd *TemplatePageData) DoAction(action string, opts ...any) string {
	log.Println("This is running")
	for _, opt := range opts {
		fmt.Printf("Option: %v\n", opt)
	}

	message := fmt.Sprintf("Requested to run action: %s\n", action)
	return message
}

func (pd *TemplatePageData) AddAction(action string, method func(args ...any), priority int32) {
	pd.actions[action] = append(pd.actions[action], Action{
		Action:   method,
		Priority: priority,
	})
}

func NewTemplatePageData() TemplatePageData {
	actions := make(map[string][]Action)
	filters := make(map[string][]Action)

	return TemplatePageData{
		actions: actions,
		filters: filters,
	}
}

func DefaultTemplateData() TemplatePageData {
	return TemplatePageData{Data: nil}
}

// Here is what I am thinking
// type PageDataWithActions interface {
// 	doAction(action string, opts ...interface{})
// 	applyFilter(filter string, opts ...interface{})
// }

func (t *Templates) Render(writer io.Writer, name string, data TemplatePageData) error {
	// FuncMap Needed?
	funcs := template.FuncMap{
		"DoAction": data.DoAction,
	}
	t.templates.Funcs(funcs)

	return t.templates.ExecuteTemplate(writer, name, data.Data)
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
