package main

import (
	"github.com/danielnill/go_hack/controllers"
	"github.com/danielnill/go_hack/web"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.FrontPageHandler).Methods("GET")
	r.HandleFunc("/discussion/{id:[0-9]+}", controllers.DiscussionHandler).Methods("GET")
	web.StartServer(r)
}
