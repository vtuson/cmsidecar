package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var rootHelmPath string

func getHelmPath(name string) string {
	if rootHelmPath == "" {
		return "/tmp/" + name
	}
	return rootHelmPath + name
}

func hasHelm() error {
	return exec.Command("helm", "--help").Run()
}
func helmPack(path string, dest string) error {
	//check folder and create if it doesnt exist
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.MkdirAll(dest, os.ModePerm)
	}
	if err := hasHelm(); err != nil {
		return errors.New("Helm is not installed")
	}

	//update dependencies
	cmd := exec.Command("helm", "dependency", "update", path)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	//pack chart
	cmd = exec.Command("helm", "package", path, "-d", dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Packaging helm charts in repo %s to %s\n", path, dest)
	return cmd.Run()
}

func helmAddRepo(name string, repo string) error {
	cmd := exec.Command("helm", "repo", "add", name, repo)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func helmPackDir(orig string, dest string) error {
	fileList := make([]string, 0)
	err := filepath.Walk(orig, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, "Chart.yaml") {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		return err
	}

	for _, file := range fileList {
		file = file[:strings.LastIndex(file, "Chart.yaml")]
		helmPack(file, dest)
	}

	return nil

}
