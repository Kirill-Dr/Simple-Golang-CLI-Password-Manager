package cmd

import (
	"fmt"
	"go-cli-password-manager/encryption"
	"go-cli-password-manager/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new password",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		key, _ := cmd.Flags().GetString("key")

		if name == "" || username == "" || password == "" || key == "" {
			fmt.Println("All fields are required!")
			return
		}

		encryptedPassword, err := encryption.Encrypt(password, key)
		if err != nil {
			fmt.Println("Error of encrypting:", err)
			return
		}

		passwords, err := storage.LoadPasswords()
		if err != nil {
			fmt.Println("Error passwords load:", err)
			return
		}

		passwords = append(passwords, storage.PasswordEntry{
			Name:     name,
			Username: username,
			Password: encryptedPassword,
		})

		err = storage.SavePasswords(passwords)
		if err != nil {
			fmt.Println("Error password save:", err)
			return
		}

		fmt.Println("Password successfully added!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("name", "", "Password name")
	addCmd.Flags().String("username", "", "User name")
	addCmd.Flags().String("password", "", "Password")
	addCmd.Flags().String("key", "", "Encrypted key")
}
