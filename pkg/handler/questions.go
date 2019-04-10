package handler

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/justina777/website-stackflow/pkg/schema"
	"github.com/justina777/website-stackflow/pkg/service"
)

const (
	MAXPAGE = 10
)

func ListHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/list.html", "templates/header.html", "templates/footer.html"))

	vPage := 1 //the page of voteLists
	cPage := 1 //the page of creationLists
	if name := r.FormValue("vp"); name != "" {
		t, err := strconv.Atoi(name)
		if err == nil {
			vPage = t
		}
	}
	if name := r.FormValue("cp"); name != "" {
		t, err := strconv.Atoi(name)
		if err == nil {
			vPage = t
		}
	}

	client := &service.StackOverflowClient{}
	voteLists := client.Fetch("votes", time.Now().AddDate(0, 0, -7), vPage)
	creationLists := client.Fetch("creation", time.Now().AddDate(0, 0, -7), cPage)

	obj := make(map[string]interface{})
	obj["VotedItems"] = TransformData(voteLists.Items)
	obj["CreationItems"] = TransformData(creationLists.Items)
	obj["VCurrentPage"] = vPage
	fmt.Println("page ", vPage, ", cpage ", cPage)

	if vPage < 3 {
		obj["VPrevPages"] = genPageNums(vPage, true, vPage-1)
		obj["VAftervPages"] = genPageNums(vPage, false, 2+3-vPage)
	} else if vPage+2 >= MAXPAGE {
		obj["VPrevPages"] = genPageNums(vPage, true, 2+2-MAXPAGE+vPage)
		obj["VAftervPages"] = genPageNums(vPage, false, MAXPAGE-vPage)
	} else {
		obj["VPrevPages"] = genPageNums(vPage, true, 2)
		obj["VAftervPages"] = genPageNums(vPage, false, 2)
	}
	fmt.Println(obj["VPrevPages"])
	fmt.Println(obj["VAftervPages"])

	if err := t.ExecuteTemplate(w, "list.html", obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func genPageNums(curPage int, isPrev bool, genNum int) []schema.Page {
	p := []schema.Page{}
	if isPrev {
		for i := genNum; i > 0; i-- {
			t := curPage - i
			item := schema.Page{
				Index: t,
			}
			p = append(p, item)
		}
	} else {
		for i := 1; i <= genNum; i++ {
			t := curPage + i
			item := schema.Page{
				Index: t,
			}
			p = append(p, item)
		}
	}

	fmt.Println(p)
	return p
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/item.html", "templates/header.html", "templates/footer.html"))

	url := r.FormValue("url")
	// url := "https://stackoverflow.com/questions/55504508/package-android-emulator-with-revision-at-least-28-1-9-not-available"

	client := &service.StackOverflowClient{}
	cont := client.GetQuestion(url)

	obj := make(map[string]interface{})
	obj["HtmlContent"] = template.HTML(html.UnescapeString(cont))

	if err := t.ExecuteTemplate(w, "item.html", obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TransformData(list []schema.Question) []schema.Question {
	for i, item := range list {
		list[i].CreationDate = time.Unix(item.IntCreationDate, 0)
		list[i].Title = html.UnescapeString(item.Title)
		list[i].Body = html.UnescapeString(item.Body)
	}
	return list
}
