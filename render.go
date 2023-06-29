package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplates(w http.ResponseWriter, gohtml string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + gohtml)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
