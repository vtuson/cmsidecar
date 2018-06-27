package main

import (
	"fmt"
	"os"
)

func main() {
	//cloneRepo("https://github.com/bitnami/charts.git", getGitPath("bitnami"))
	if err := helmAddRepo("bitnami", "https://charts.bitnami.com/bitnami"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	helmPackDir(getGitPath("bitnami"), getHelmPath("helm"))
}
