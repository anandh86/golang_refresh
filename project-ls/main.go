package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"strings"
	"flag"
)

func main() {
	// Your solution goes here. Good luck!
	var showHiddenFiles = flag.Bool("a", false, "show hidden files")

	flag.Parse()

	files := listFiles("testdata")

	for _, file := range files {

		if *showHiddenFiles || !strings.HasPrefix(file, ".") {
			fmt.Println(file)
		}
	}

	//singleLine := strings.Join(files, " ")

	//fmt.Println(singleLine)
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
