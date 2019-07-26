package http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GorillaServer() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}/code/{code}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["name"]
		page := vars["code"]

		_, _ = fmt.Fprintf(writer, "You've requested user: %s with code %s \n", title, page)

	}).Methods("POST").Host("localhost")

	// subRoute
	userRoute := r.PathPrefix("/users").Subrouter()
	userRoute.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Welcome gorilla endpoint")
	})

	userRoute.HandleFunc("/insecure", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "insecure route")
	})

	_ = http.ListenAndServe(":8080", r)
}