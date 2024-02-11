package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GenerateCmd() {
	rootCmd := &cobra.Command{
		Use:   "locker",
		Short: "Cmd based password locker",
	}

	rootCmd.AddCommand(encryptCmd(), decryptCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func encryptCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "convert the normal file into encrypted file",
		Run:   encryptFile,
	}

	cmd.Flags().StringP("file", "f", "", "Enter the file path")
	cmd.Flags().StringP("password", "p", "", "Password")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("password")

	return cmd
}

func decryptCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "convert the encrypted file into normal file",
		Run:   decryptFile,
	}

	cmd.Flags().StringP("file", "f", "", "Enter the file path")
	cmd.Flags().StringP("password", "p", "", "Password")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("password")

	return cmd
}
