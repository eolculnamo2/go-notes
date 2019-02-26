package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`
		<h1>This is multiline</h1>
		<p>New line</p>
		`)
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<h1>You %s even pass %s</h1>", "can", "variables")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}