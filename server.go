package main 

import ("html/template",
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){

	t, _ :=template.ParseFiles("index.html")
}

func main() {
	http.HandleFunc("/", index_handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}