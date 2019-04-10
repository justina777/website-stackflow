package main

import (
	"fmt"
	"net/http"

	"github.com/justina777/website-stackflow/pkg/handler"
)

func init() {
	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
}

//Go application entrypoint
func main() {

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/sorry", handler.SorryHandler)
	http.HandleFunc("/list", handler.ListHandler)

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	http.ListenAndServe(":8080", nil)
}
