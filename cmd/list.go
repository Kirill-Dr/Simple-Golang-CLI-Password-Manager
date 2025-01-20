package cmd

import (
	"fmt"
	"go-cli-password-manager/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all passwords list",
	Run: func(cmd *cobra.Command, args []string) {
		passwords, err := storage.LoadPasswords()
		if err != nil {
			fmt.Println("Error of loading passwords:", err)
			return
		}

		if len(passwords) == 0 {
			fmt.Println("No saved passwords.")
			return
		}

		for _, entry := range passwords {
			fmt.Printf("Password name: %s, User name: %s\n", entry.Name, entry.Username)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
