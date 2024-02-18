package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gerlacdt/gcurl/http"
)

var rootCmd = &cobra.Command{
	Use:   "gcurl",
	Short: "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Long:  "gcurl is a replacement for curl to make HTTP requests from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		http.Get(url)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
