package handler

import (
	"html"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/justina777/website-stackflow/pkg/schema"
	"github.com/justina777/website-stackflow/pkg/service"
	"github.com/justina777/website-stackflow/pkg/tool"
	cache "github.com/patrickmn/go-cache"
)

const (
	MAXPAGE   = 10
	PAGE_SIZE = 10
)

var (
	c *cache.Cache
)

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func findHotTags(sort string, key string) schema.PairList {

	tList, found := c.Get(key)
	if found {
		return tList.(schema.PairList)
	}

	client := &service.StackOverflowClient{}
	list := client.Fetch(sort, time.Now().AddDate(0, 0, -7), 1, 100)

	mapTags := make(map[string]int)
	for _, item := range list.Items {
		for _, t := range item.Tags {
			if v, ok := mapTags[t]; ok {
				mapTags[t] = v + 1
			} else {
				mapTags[t] = 1
			}
		}
	}

	tList = sortMapByValue(mapTags)

	c.Set(key, tList, cache.DefaultExpiration)
	return tList.(schema.PairList)
}

func sortMapByValue(m map[string]int) schema.PairList {
	p := make(schema.PairList, len(m))
	i := 0
	for k, v := range m {
		if v > 1 {
			p[i] = schema.Pair{k, v}
			i++
		}
	}
	sort.Sort(p)
	if i > 11 {
		return p[1:11]
	} else {
		return p[1:]
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/list.html", "templates/list-newest.html", "templates/header.html", "templates/footer.html"))

	pageType := ""
	pageNum := 1 //the page of voteLists
	if name := r.FormValue("vp"); name != "" {
		t, err := strconv.Atoi(name)
		if err == nil {
			pageNum = t
		}
	}
	if name := r.FormValue("t"); name != "" {
		pageType = name
	}

	client := &service.StackOverflowClient{}

	obj := make(map[string]interface{})
	if pageType == "" {
		hotTags := findHotTags("votes", "voteTags")
		voteLists := client.Fetch("votes", time.Now().AddDate(0, 0, -7), pageNum, PAGE_SIZE)
		obj["VotedItems"] = TransformData(voteLists.Items)
		obj["HotTags"] = hotTags
	} else {
		creationLists := client.Fetch("creation", time.Now().AddDate(0, 0, -7), pageNum, PAGE_SIZE)
		hotTags := findHotTags("creation", "creationTags")
		obj["CreationItems"] = TransformData(creationLists.Items)
		obj["HotTags"] = hotTags
	}

	obj["Logged"] = tool.IsLogin(r.FormValue("l"))
	obj["CurrentPage"] = pageNum
	if pageNum < 3 {
		obj["PrevPages"] = genPageNums(pageNum, true, pageNum-1)
		obj["AftervPages"] = genPageNums(pageNum, false, 2+3-pageNum)
	} else if pageNum+2 >= MAXPAGE {
		obj["PrevPages"] = genPageNums(pageNum, true, 2+2-MAXPAGE+pageNum)
		obj["AftervPages"] = genPageNums(pageNum, false, MAXPAGE-pageNum)
	} else {
		obj["PrevPages"] = genPageNums(pageNum, true, 2)
		obj["AftervPages"] = genPageNums(pageNum, false, 2)
	}
	// fmt.Println("page ", vPage, ", cpage ", cPage)

	if pageType == "" {
		if err := t.ExecuteTemplate(w, "list.html", obj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := t.ExecuteTemplate(w, "list-newest.html", obj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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

	// fmt.Println(p)
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
		list[i].Owner.DisplayName = html.UnescapeString(item.Owner.DisplayName)
		timeDiff := time.Now().UTC().Sub(list[i].CreationDate)
		if timeDiff.Hours() >= 24 {
			list[i].PostedTimeAgo = int(timeDiff.Hours()) / 24
			list[i].UnitPeroid = "days"
		} else if timeDiff.Hours() > 1 {
			list[i].PostedTimeAgo = int(timeDiff.Hours())
			list[i].UnitPeroid = "hours"
		} else if timeDiff.Minutes() > 1 {
			list[i].PostedTimeAgo = int(timeDiff.Minutes())
			list[i].UnitPeroid = "mins"
		} else {
			list[i].PostedTimeAgo = int(timeDiff.Seconds())
			list[i].UnitPeroid = "secs"
		}
	}
	return list
}
