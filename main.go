package main

import (
	"net/http"
	"new_practice/awesomeProject/server"
)

func main() {

	http.HandleFunc("/valid", server.Valid)
	http.HandleFunc("/link/", server.Link)
	http.ListenAndServe(":8090", nil)
}
