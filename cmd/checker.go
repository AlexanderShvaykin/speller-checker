package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"ya-speller/internal/httpclient"
)

func checkFiles(path string) {

	files, err := readFiles(path)
	if err != nil {
		log.Fatal("Dir not found")
	}
	var osFiles []*os.File
	defer func() {
		for _, file := range osFiles {
			file.Close()
		}
	}()
	for _, file := range files {
		file, err := os.Open(file.root + "/" + file.name)
		osFiles = append(osFiles, file)

		if err != nil {
			log.Fatal(err)
		}

		urls := make(chan string)
		go func(out chan<- string, f *os.File) {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				text := strings.TrimSpace(scanner.Text())
				url := httpclient.BuildURL(baseURL, map[string]string{"text": `'` + text + `'`})
				urls <- url
			}
			close(out)
		}(urls, file)

		for u := range urls {
			in := make(chan Mistake)
			go func(out chan<- Mistake, u string) {
				var mistakes []Mistake
				response := httpclient.Get(u)
				if err := json.Unmarshal([]byte(response), &mistakes); err != nil {
					fmt.Println(err, u, response)
				}

				for _, m := range mistakes {
					out <- m
				}
				close(out)
			}(in, u)

			for m := range in {
				exitCode = 1
				fmt.Println(m)
			}
		}
	}
	os.Exit(exitCode)
}
