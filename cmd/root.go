package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-password-manager",
	Short: "CLI password manager",
	Long:  `This is a simple CLI application for password management.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use commands: add, list, delete, update, find.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
