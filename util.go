package main

import (
	"io/ioutil"
	"os"
	"path"
)

func readDir(dirname string) (ret []*os.FileInfo) {
	ret, err := ioutil.ReadDir(dirname)
	if err != nil {
		// the directory doesn't exist - create it
		err = os.MkdirAll(dirname, 0755)
	}
	return
}

func walkDir(dirname string, v path.Visitor) {
	// first make sure the directory exists
	readDir(dirname)
	// now walk it
	path.Walk(dirname, v, nil)
}

func readFile(filename string) (contents string) {
	contentarry, err := ioutil.ReadFile(filename)
	if err != nil {
		// the file doesn't exist - create the directory and the file
		dirname, _ := path.Split(filename)
		err = os.MkdirAll(dirname, 0755)
		ioutil.WriteFile(filename, []byte{}, 0644)
	}
	contents = string(contentarry)
	return
}
