package tree

import (
	"os"
	"path/filepath"
)

func filesList(root string) ([]string, error) {
	var files []string
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return files, err
	} else {
		err = nil
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
