package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerReady(w http.ResponseWriter, req *http.Request) {
	response(w, 200, "ready")
}

//name and addr for adding a repo
func handlerAddHelmRepo(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	addr := req.FormValue("addr")
	if err := helmAddRepo(name, addr); err != nil {
		response(w, 400, "Failed to add repo: "+addr)
		return
	}
	response(w, 200, "ready")
}

//name and addr for adding a repo
func handlerNewHelmRepo(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	git := req.FormValue("git")

	if git == "" {
		response(w, 400, "Please provide public git endpoint")
	}

	if err := cloneRepo(git, getGitPath(name)); err != nil {
		response(w, 400, "Failed to clone repo: "+git)
		return
	}
	go helmPackDir(getGitPath(name), getHelmPath(name))
	response(w, 200, "syncing and packing repo")
}

//name and addr for adding a repo
func handlerUpdateHelmRepo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if err := pullRepo(getGitPath(vars["name"])); err != nil {
		response(w, 400, "Failed to clone repo: "+vars["name"])
		return
	}
	go helmPackDir(getGitPath(vars["name"]), getHelmPath(vars["name"]))
	response(w, 200, "syncing and packing repo")
}

func response(w http.ResponseWriter, code int, value string) {
	w.WriteHeader(code)
	fmt.Fprintln(w, value)
}
