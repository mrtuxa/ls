package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(getDirectory())
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		colorRed := "\033[31m"
		colorReset := "\033[0m"

		fmt.Println(string(colorRed), f.Name())
		fmt.Println(string(colorReset))
	}
}

func getDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return path
}
