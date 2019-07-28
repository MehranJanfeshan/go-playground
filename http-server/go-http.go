package http_server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type payload struct {
	FirstName string
	LastName  string
	Email     string
}

func GoHttpServer() {
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)

		if err != nil {
			panic(err)
		}

		var t payload
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		log.Println(t)

		_, _ = fmt.Fprintf(writer, "welcome to the website your first name is: %s and your last name is: %s", t.FirstName, t.LastName)
	})

	// file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(":80", nil)
}
