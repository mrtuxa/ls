package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		listDir()
	case "linux":
		fmt.Println("Linux.")
		listDir()
	case "windows":
		log.Fatal("\nWindows is currently not supported")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
		listDir()
	}
}

func getDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return path
}

func listDir() {
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
