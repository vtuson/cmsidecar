package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()
	flag.StringVar(&rootGitPath, "git", "/tmp/", "specify path to store git clone repos")
	flag.StringVar(&rootHelmPath, "helm", "/tmp/", "specify path to store helm packs")
	flag.Parse()
	//routes
	r.HandleFunc("/", handlerReady).Methods("GET")
	r.HandleFunc("/repo/new", handlerNewHelmRepo).Methods("POST")
	r.HandleFunc("/repo/dependency", handlerAddHelmRepo).Methods("POST")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	http.ListenAndServe(":3000", n)
}
