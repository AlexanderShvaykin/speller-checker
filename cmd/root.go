package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"ya-speller/api"

	"github.com/spf13/cobra"
)

var path string

const baseURL string = "https://speller.yandex.net/services/spellservice.json/checkText"

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

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()
				url := api.BuildURL(baseURL, map[string]string{"text": string(line)})
				var mistakes []Mistake
				if err := json.Unmarshal([]byte(api.Get(url)), &mistakes); err != nil {
					log.Fatal(err)
				}
				if len(mistakes) > 0 {
					fmt.Println(mistakes)
				}
			}
		}
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
}
