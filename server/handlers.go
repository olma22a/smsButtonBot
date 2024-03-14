package server

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

var m = make(map[string]string)

func Valid(w http.ResponseWriter, req *http.Request) {
	link := req.URL.Query().Get("link") // чтобы изымать по определнному парамтеру

	id := uuid.New() //generim uniq id

	_, err := url.ParseRequestURI(link)
	if err != nil {
		fmt.Println(" не норм ", err)
	} else {
		m[id.String()] = link
		fmt.Fprintf(w, "id: %s ", id.String())

	}

}

func Link(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	validLink := m[id]
	if validLink != "" {

		fmt.Fprintf(w, "%s\n", validLink)

	} else {

		fmt.Fprintf(w, "нету\n")
	}

}
