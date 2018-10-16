package main 

import ("html/template",
		"net/http"

		"google.golang.org/appengine")

func main() {
        http.HandleFunc("/", handle)
        appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
	            http.Redirect(w, r, "/", http.StatusFound)
	            return
	        }
	    fmt.Fprintln(w, "Hello, world!")
}