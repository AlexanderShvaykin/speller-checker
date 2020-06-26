package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ya-speller/internal/checker"
)

var path string

const baseURL string = "https://speller.yandex.net/services/spellservice.json/checkText"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ya-speller",
	Short: "Yandex speller checker",
	Run: func(cmd *cobra.Command, _args []string) {
		os.Exit(checker.CheckFiles(path, baseURL))
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
