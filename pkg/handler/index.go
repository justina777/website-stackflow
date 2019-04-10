package handler

import (
	"github.com/justina777/website-stackflow/pkg/schema"
	"html/template"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := schema.Header{"Android template", time.Now().Format(time.Stamp)}

	t := template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html"))
	//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
	// if name := r.FormValue("name"); name != "" {
	// 	header.Name = name
	// }
	//If errors show an internal server error message
	//I also pass the welcome struct to the welcome-template.html file.
	if err := t.ExecuteTemplate(w, "index.html", header); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
