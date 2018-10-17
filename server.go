package main 

import (
		"net/http"
		"html/template"
		"fmt"

		"google.golang.org/appengine"
		)
		
		var (
			indexTemplate = template.Must(template.ParseFiles("index.gtpl"))
			)

		type templateParams struct {
				Password string
				Name     string
				Notice   string
		}

func main() {
	    http.HandleFunc("/", handle)
		appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
				http.Redirect(w, r, "/", http.StatusFound)
				return
		}

		params := templateParams{}
			
		if r.Method == "GET" {
			indexTemplate.Execute(w, params)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		params.Name = username
		params.Password = password
		if username == "ashlynn" {
			if password == "password" {
				params.Notice = fmt.Sprintf("Successfully logged in\nUsername: %s\nPassword: %s", username, password)
			}
		} else {
			params.Notice = fmt.Sprintf("Ah ah ah, you didn't say the magic word!\nUsername: %s\nPassword: %s", username, password)
		}

		indexTemplate.Execute(w, params)

}