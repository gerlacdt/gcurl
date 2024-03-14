package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gerlacdt/gcurl/http"
)

var method string
var headers []string
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "gcurl",
	Short: "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Long:  "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		givenUrl := args[0]

		switch method {
		case "GET":
			{
				response, err := http.Get(givenUrl, verbose)
				if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
				}
				response.Print(verbose)
			}
		case "POST":
			{
				response, err := http.Post(givenUrl, headers, os.Stdin)
				if err != nil {
					fmt.Printf("%s", err)
					os.Exit(1)
				}
				response.Print(verbose)
			}
		default:
			{
				fmt.Printf("Invalid HTTP method given, got: %s", method)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&method, "method", "X", "GET", "http method to use")
	slice := make([]string, 0)
	rootCmd.PersistentFlags().StringSliceVarP(&headers, "header", "H", slice, "http header to put in request")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose flag, print out http headers for request and response")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
