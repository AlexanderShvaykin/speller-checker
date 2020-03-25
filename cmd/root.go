package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"ya-speller/internal/httpclient"

	"github.com/spf13/cobra"
)

var path string

const baseURL string = "https://speller.yandex.net/services/spellservice.json/checkText"

var exitCode int

// Mistake is spelling mistake
type Mistake struct {
	Word string
	S    []string
}

func (m Mistake) String() string {
	return fmt.Sprintf("[word: %s, suggestions: %s]", m.Word, m.S)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ya-speller",
	Short: "Yandex speller checker",
	Run: func(cmd *cobra.Command, _args []string) {

		files, err := readFiles(path)
		if err != nil {
			log.Fatal("Dir not found")
		}
		for _, file := range files {
			file, err := os.Open(file.root + "/" + file.name)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			urls := make(chan string)
			go func(out chan<- string, f *os.File) {
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					text := strings.TrimSpace(scanner.Text())
					url := httpclient.BuildURL(baseURL, map[string]string{"text": `'` + string(text) + `'`})
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

	}, // end Run
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Path to dir")
}
