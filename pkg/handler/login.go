package handler

import (
	"github.com/justina777/website-stackflow/pkg/tool"
	"fmt"
	"net/http"
)

var (
	authUserName = []string{"heng_li", "justina_lin"}
	authPassword = []string{"thisispassword", "password9876", "1234aaa"}
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// u := dao.AuthUser{
	// 	UserName: "Android template",
	// 	Password: time.Now().Format(time.Stamp),
	// }

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		r.ParseForm()
		// logic part of log in
		u := r.Form["username"]
		p := r.Form["password"]

		fmt.Println("user: ", u, ", password:", p)
		if len(u) == 0 || len(p) == 0 {
			http.Redirect(w, r, "/sorry", http.StatusSeeOther)
		}

		fmt.Println("user found:", tool.Find(authUserName, u[0]), ", passowrd found:", tool.Find(authPassword, p[0]))
		if tool.Find(authUserName, u[0]) >= 0 && tool.Find(authPassword, p[0]) >= 0 {
			http.Redirect(w, r, "/list", http.StatusSeeOther)

		} else {
			http.Redirect(w, r, "/sorry", http.StatusSeeOther)
		}
	} else {
		fmt.Println("sorry page")
		http.Redirect(w, r, "/sorry", http.StatusSeeOther)
	}

}
