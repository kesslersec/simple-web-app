package main 

import (
		"net/http"
		)

func main() {
        http.HandleFunc("/", handle)
        if err := http.ListenAndServe(":8080",nil); err != nil {
        	panic(err)
        }

		appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
	            http.Redirect(w, r, "/", http.StatusFound)
	            return
	        }
	    w.Write([]byte("weow"))
}