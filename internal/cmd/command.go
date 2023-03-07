package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-line-bot",
	Short: "This is the first go-line-bot command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Port string
var getPortCmd = &cobra.Command{
	Use:   "port",
	Short: "port",
}

func init() {
	rootCmd.AddCommand(getPortCmd)
	getPortCmd.Flags().StringVar(&Port, "set", "8080", "input your port")
}
