package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var rootPath string

func hasGit() error {
	return exec.Command("git", "--version").Run()
}

func getPath(name string) string {
	if rootPath == "" {
		return "/tmp/" + name
	}
	return rootPath + name
}

func hasRepo(name string) bool {
	if _, err := os.Stat(getPath(name)); err == nil {
		return false
	}
	return true
}

func cloneRepo(orig string, dest string) error {
	if err := hasGit(); err != nil {
		fmt.Println(err)
		return errors.New("Git is not installed")
	}
	cmd := exec.Command("git", "clone", orig, dest)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func pullRepo(path string) error {
	if err := hasGit(); err != nil {
		fmt.Println(err)
		return errors.New("Git is not installed")
	}
	cmd := exec.Command("git", "pull", path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
