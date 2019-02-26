package main

import ("fmt"
		"net/http"
		"html/template")

func index_handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "derp")
	t,_ := template.ParseFiles("template.html")
	fmt.Println("template")
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}