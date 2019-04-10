package handler

import (
	"html"
	"html/template"
	"net/http"
	"time"

	"github.com/justina777/website-stackflow/pkg/schema"
	"github.com/justina777/website-stackflow/pkg/service"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/list.html", "templates/header.html", "templates/footer.html"))

	// if name := r.FormValue("name"); name != "" {
	// 	header.Name = name
	// }

	client := &service.StackOverflowClient{}
	voteLists := client.Fetch("votes", time.Now().AddDate(0, 0, -7))
	creationLists := client.Fetch("creation", time.Now().AddDate(0, 0, -7))

	obj := make(map[string]interface{})
	obj["VotedItems"] = TransformData(voteLists.Items)
	obj["CreationItems"] = TransformData(creationLists.Items)

	if err := t.ExecuteTemplate(w, "list.html", obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TransformData(list []schema.Question) []schema.Question {
	for i, item := range list {
		list[i].CreationDate = time.Unix(item.IntCreationDate, 0)
		list[i].Title = html.UnescapeString(item.Title)
	}
	return list
}
