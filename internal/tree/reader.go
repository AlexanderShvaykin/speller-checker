package tree

import (
	"bufio"
	"os"
	"strings"
)

func ReadFiles(root string, handler func(line string)) error {
	fileNames, err := filesList(root)
	if err == nil {
		for _, path := range fileNames {
			file, err := os.Open(path)
			if err == nil {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					text := strings.TrimSpace(scanner.Text())
					handler(text)
				}
			}
		}
	}
	return err
}
