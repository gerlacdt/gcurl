package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gerlacdt/gcurl/http"
)

var method string

var rootCmd = &cobra.Command{
	Use:   "gcurl",
	Short: "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Long:  "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must provide exactly one argument for the URL")
			os.Exit(1)
		}
		givenUrl := args[0]
		body, err := http.Get(givenUrl)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s", body)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&method, "method", "X", "GET", "http method to use")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
