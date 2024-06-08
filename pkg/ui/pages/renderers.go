package pages

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	templatesBasePath = "./configs/templates"
)

func GoTemplateRender(w http.ResponseWriter, _ *http.Request, tmpl string, data map[string]any, partials ...string) {
	tplList := []string{
		fmt.Sprintf("%s/%s.gotmpl", templatesBasePath, tmpl),
	}
	if len(partials) > 0 {
		for _, item := range partials {
			tplList = append(tplList, fmt.Sprintf("%s/%s.gotmpl", templatesBasePath, item))
		}
	}

	tpl, err := template.ParseFiles(tplList...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
