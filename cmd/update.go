package cmd

import (
	"fmt"
	"go-cli-password-manager/encryption"
	"go-cli-password-manager/storage"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update password for resource",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		newPassword, _ := cmd.Flags().GetString("password")
		key, _ := cmd.Flags().GetString("key")

		if name == "" || newPassword == "" || key == "" {
			fmt.Println("All fields are required!")
			return
		}

		encryptedPassword, err := encryption.Encrypt(newPassword, key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		passwords, err := storage.LoadPasswords()
		if err != nil {
			fmt.Println("Error loading passwords:", err)
			return
		}

		var found bool
		for i, entry := range passwords {
			if entry.Name == name {
				passwords[i].Password = encryptedPassword
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Password with this name not found.")
			return
		}

		err = storage.SavePasswords(passwords)
		if err != nil {
			fmt.Println("Error saving password:", err)
			return
		}

		fmt.Println("Password successfully updated!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("name", "", "Password name to update")
	updateCmd.Flags().String("password", "", "New password")
	updateCmd.Flags().String("key", "", "Encryption key")
}
