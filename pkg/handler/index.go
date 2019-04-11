package handler

import (
	"html/template"
	"net/http"

	"github.com/justina777/website-stackflow/pkg/tool"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html"))

	// if name := r.FormValue("name"); name != "" {
	// 	header.Name = name
	// }
	obj := make(map[string]interface{})
	obj["Logged"] = tool.IsLogin(r.FormValue("l"))

	if err := t.ExecuteTemplate(w, "index.html", obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
