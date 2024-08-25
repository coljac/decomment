package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type FileLister struct{}

func (fl *FileLister) ListFiles() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	fl := &FileLister{}
	fl.ListFiles()
}

