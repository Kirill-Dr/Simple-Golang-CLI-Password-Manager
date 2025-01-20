package cmd

import (
	"fmt"
	"go-cli-password-manager/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove saved password",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
			fmt.Println("Password name is required!")
			return
		}

		passwords, err := storage.LoadPasswords()
		if err != nil {
			fmt.Println("Error of loading passwords:", err)
			return
		}

		var indexToRemove int = -1
		for i, entry := range passwords {
			if entry.Name == name {
				indexToRemove = i
				break
			}
		}

		if indexToRemove == -1 {
			fmt.Println("Password with this name not found.")
			return
		}

		passwords = append(passwords[:indexToRemove], passwords[indexToRemove+1:]...)

		err = storage.SavePasswords(passwords)
		if err != nil {
			fmt.Println("Error saving data:", err)
			return
		}

		fmt.Println("Password successfully removed!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("name", "", "Password name to delete")
}
