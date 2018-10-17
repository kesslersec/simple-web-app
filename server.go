package main 

import (
		"net/http"
		"html/template"
		"fmt"
		"bytes"

		"google.golang.org/appengine"
		)
		
		var (
			indexTemplate = template.Must(template.ParseFiles("index.gtpl"))
			)

		type templateParams struct {
				Password string
				Name     string
				Notice   string
				Html	 template.HTML
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

		var buffer bytes.Buffer
		username := r.FormValue("username")
		password := r.FormValue("password")
		params.Name = username
		params.Password = password
		if username == "ashlynn" {
			if password == "password" {
				buffer.WriteString("Successfully logged in")
			}
		} else {
			buffer.WriteString("Ah ah ah, you didn't say the magic word!")
		}
		if(username != "" && password != ""){
			params.Html == fmt.Sprintf("<br>Username: %s<br>Password: %s", username, password)
		}
		params.Notice = buffer.String()

		indexTemplate.Execute(w, params)
}
