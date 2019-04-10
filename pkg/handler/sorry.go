package handler

import (
	"html/template"
	"net/http"
)

func SorryHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/sorry.html", "templates/header.html", "templates/footer.html"))
	if err := t.ExecuteTemplate(w, "sorry.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
