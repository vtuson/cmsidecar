package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	help := false
	r := mux.NewRouter()
	flag.StringVar(&rootGitPath, "git", "/tmp/", "specify path to store git clone repos")
	flag.StringVar(&rootHelmPath, "helm", "/tmp/", "specify path to store helm packs")
	flag.BoolVar(&help, "help", false, "print help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	//routes
	r.HandleFunc("/", handlerReady).Methods("GET")
	r.HandleFunc("/repo/new", handlerNewHelmRepo).Methods("POST")
	r.HandleFunc("/repo/dependency", handlerAddHelmRepo).Methods("POST")
	r.HandleFunc("/repo/{name}/update", handlerUpdateHelmRepo).Methods("GET")
	r.HandleFunc("/repo/{name}", handlerReady).Methods("DELETE")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	http.ListenAndServe(":3000", n)
}
