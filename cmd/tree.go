package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
)

type treeFile struct {
	root      string
	name      string
	isDir     bool
	childrens []treeFile
	size      int64
}

func (f *treeFile) readChildrens() bool {
	if f.isDir {
		childrens, err := readFiles(f.root + "/" + f.name)
		if err == nil {
			f.childrens = childrens
			return true
		}

		log.Fatalf("Error: %v", err)
		return false
	}
	return true
}

func (f treeFile) String() string {
	return fmt.Sprintf("%v size %v | ", f.name, f.size)
}

func readFiles(root string) ([]treeFile, error) {
	var files []treeFile
	fileInfo, err := ioutil.ReadDir(root)

	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		myFile := treeFile{name: file.Name(), isDir: file.IsDir(), root: root, size: file.Size()}
		myFile.readChildrens()
		files = append(files, myFile)
	}
	return files, nil
}
