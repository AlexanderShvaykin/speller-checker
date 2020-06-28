package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ya-speller/internal/httpclient"
	"ya-speller/internal/tree"
)

var (
	path    string
	verbose bool
)

const baseURL string = "https://speller.yandex.net/services/spellservice.json/checkText"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ya-speller",
	Short: "Yandex speller checker",
	Run: func(cmd *cobra.Command, _args []string) {
		mistakeChan := make(chan Mistake, 0)
		successChan := make(chan string, 0)
		exitCode := 0
		lineCount, err := tree.ReadFiles(path, func(line string) {
			go func(text string, out chan<- Mistake, success chan<- string) {
				url := httpclient.BuildURL(baseURL, map[string]string{"text": `'` + text + `'`})
				var mistakes []Mistake
				response := httpclient.Get(url)
				if err := json.Unmarshal([]byte(response), &mistakes); err != nil {
					fmt.Println(err, url, response)
				}
				if len(mistakes) > 0 {
					for _, m := range mistakes {
						out <- m
					}
				} else {
					success <- text
				}

			}(line, mistakeChan, successChan)
		})

		if err != nil {
			panic(err)
		}

		for i := 0; i < lineCount; i++ {
			select {
			case mistake := <-mistakeChan:
				exitCode = 1
				fmt.Println(mistake)
			case text := <-successChan:
				if verbose {
					fmt.Printf("%v -- OK\n", text)
				}
			}
		}
		os.Exit(exitCode)
	},
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
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode")
}
