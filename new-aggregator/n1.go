package main

import ("fmt"
		"net/http"
		"io/ioutil")

func main() {
	//_ means its throw away that
	// you dont intend to use
	resp, _ := http.Get("https://www.sword-point.com/sitemap")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close() // frees up resources from request
}