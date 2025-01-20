package cmd

import (
	"fmt"
	"go-cli-password-manager/encryption"
	"go-cli-password-manager/storage"

	"github.com/spf13/cobra"
)

var findPasswordDetailCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a password by resource name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must provide the resource name.")
			return
		}

		name := args[0]
		key, _ := cmd.Flags().GetString("key")
		if key == "" {
			fmt.Println("Encryption key is required!")
			return
		}

		passwords, err := storage.LoadPasswords()
		if err != nil {
			fmt.Println("Error loading passwords:", err)
			return
		}

		var foundPassword *storage.PasswordEntry
		for _, password := range passwords {
			if password.Name == name {
				foundPassword = &password
				break
			}
		}

		if foundPassword == nil {
			fmt.Println("Password for the resource not found.")
			return
		}

		decryptedPassword, err := encryption.Decrypt(foundPassword.Password, key)
		if err != nil {
			fmt.Println("Error decrypting the password:", err)
			return
		}

		fmt.Printf("Password Name: %s\n", foundPassword.Name)
		fmt.Printf("Username: %s\n", foundPassword.Username)
		fmt.Printf("Password: %s\n", decryptedPassword)
	},
}

func init() {
	rootCmd.AddCommand(findPasswordDetailCmd)
	findPasswordDetailCmd.Flags().String("key", "", "Encryption key")
}
