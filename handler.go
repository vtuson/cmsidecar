package main

import (
	"fmt"
	"net/http"
	"os"

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

//name of repo to delete
//it needs to delete the git repo and the helm folder
func handlerDeleteRepo(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if err := os.RemoveAll(getGitPath(name)); err != nil {
		response(w, 500, "could not delete git repo:"+name)
		return
	}
	if err := os.RemoveAll(getHelmPath(name)); err != nil {
		response(w, 500, "could not delete git repo:"+name)
		return
	}
	response(w, 200, "OK")

}

func response(w http.ResponseWriter, code int, value string) {
	w.WriteHeader(code)
	fmt.Fprintln(w, value)
}
