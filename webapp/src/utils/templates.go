package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template
 
// CarregarTemplate insere os templates html na variável templates
func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}