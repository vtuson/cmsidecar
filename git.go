package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var rootGitPath string

func hasGit() error {
	return exec.Command("git", "--version").Run()
}

func getGitPath(name string) string {
	if rootGitPath == "" {
		return "/tmp/" + name
	}
	return rootGitPath + name
}

func hasRepo(name string) bool {
	if _, err := os.Stat(getGitPath(name)); err == nil {
		return false
	}
	return true
}

func cloneRepo(orig string, dest string) error {
	if err := hasGit(); err != nil {
		fmt.Println(err)
		return errors.New("Git is not installed")
	}
	os.RemoveAll(dest)
	cmd := exec.Command("git", "clone", orig, dest)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Printf("Cloning %s into %s\n", orig, dest)
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
	fmt.Printf("Pulling repo in %s\n", path)
	return cmd.Run()
}
