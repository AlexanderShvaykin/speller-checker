package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ya-speller/api"

	"github.com/spf13/cobra"
)

var path string

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
				fmt.Println(line)
				url := "https://speller.yandex.net/services/spellservice.json/checkText?text=" + string(line)

				fmt.Println("Body:", api.Get(url))
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
