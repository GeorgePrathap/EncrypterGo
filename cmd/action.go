package cmd

import (
	"encrypterGo/utils"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Encrypt struct {
	Source   string
	Target   string
	Password string
}

func GetFlagValue(cmd *cobra.Command, flag string) string {
	value, _ := cmd.Flags().GetString(flag)
	return value
}

func (e *Encrypt) validation() error {
	if e.Source == "" {
		return errors.New("Source path should not be empty")
	}

	if e.Password == "" {
		return errors.New("Password should not be empty")
	}

	_, err := os.Stat(e.Source)
	if err != nil {
		return err
	}

	return nil
}

func processFile(cmd *cobra.Command, isEncryption bool) (fileName string, err error) {
	decrypt := Encrypt{
		Source:   GetFlagValue(cmd, "file"),
		Password: GetFlagValue(cmd, "password"),
	}

	if err = decrypt.validation(); err != nil {
		return
	}

	buffer, err := utils.ReadFromFile(decrypt.Source)
	if err != nil {
		return
	}

	fileInfo, err := utils.GetFileStat(decrypt.Source)
	if err != nil {
		return
	}

	fileName = utils.ChangeFileExtension(fileInfo.Name(), ".txt")
	if isEncryption {
		fileName = utils.ChangeFileExtension(fileInfo.Name(), ".aed")
	}

	file, err := utils.CreateFile(fileName)
	if err != nil {
		return
	}

	hash := utils.GenerateTextToHash(decrypt.Password)

	if isEncryption {
		encryptedText := utils.Encrypt(fmt.Sprintf("%s", &buffer), hash)
		err = utils.WriteToFile(file, encryptedText)
		if err != nil {
			return
		}
	} else {
		decryptedText := utils.Decrypt(buffer, hash)
		err = utils.WriteToFile(file, []byte(decryptedText))
		if err != nil {
			return
		}
	}

	return
}

func decryptFile(cmd *cobra.Command, args []string) {
	fileName, err := processFile(cmd, false)
	if err != nil {
		fmt.Printf("Error : %v\n", err.Error())
		return
	}

	fmt.Printf("File created successfully : %v\n", fileName)
}

func encryptFile(cmd *cobra.Command, args []string) {
	fileName, err := processFile(cmd, true)
	if err != nil {
		fmt.Printf("Error : %v\n", err.Error())
		return
	}

	fmt.Printf("File created successfully : %v\n", fileName)
}

// func decryptFile(cmd *cobra.Command, args []string) {
// 	decrypt := Encrypt{
// 		Source:   GetFlagValue(cmd, "file"),
// 		Password: GetFlagValue(cmd, "password"),
// 	}

// 	buffer, err := utils.ReadFromFile(decrypt.Source)
// 	if err != nil {
// 		fmt.Printf("There is no file in the specified location: %v\n", decrypt.Source)
// 		return
// 	}

// 	fileInfo, err := utils.GetFileStat(decrypt.Source)
// 	if err != nil {
// 		fmt.Printf("Error in getting the file state: %v\n", err.Error())
// 		return
// 	}

// 	newFileName := utils.ChangeFileExtension(fileInfo.Name(), ".txt")
// 	file, err := utils.CreateFile(newFileName)
// 	if err != nil {
// 		fmt.Printf("Error in creating the file : %v\n", err.Error())
// 		return
// 	}

// 	hash := utils.GenerateTextToHash(decrypt.Password)

// 	decryptedText := utils.Decrypt(buffer, hash)
// 	err = utils.WriteToFile(file, []byte(decryptedText))
// 	if err != nil {
// 		fmt.Printf("Error in writing the file: %v\n", err.Error())
// 		return
// 	}

// 	fmt.Printf("converted successfully file name is : %v\n", file.Name())

// }

// func encryptFile(cmd *cobra.Command, args []string) {
// 	encrypt := Encrypt{
// 		Source:   GetFlagValue(cmd, "file"),
// 		Password: GetFlagValue(cmd, "password"),
// 	}

// 	buffer, err := utils.ReadFromFile(encrypt.Source)
// 	if err != nil {
// 		fmt.Printf("There is no file in the specified location: %v\n", encrypt.Source)
// 		return
// 	}

// 	fileInfo, err := utils.GetFileStat(encrypt.Source)
// 	if err != nil {
// 		fmt.Printf("Error in getting the file state: %v\n", err.Error())
// 		return
// 	}

// 	newFileName := utils.ChangeFileExtension(fileInfo.Name(), ".aed")
// 	file, err := utils.CreateFile(newFileName)
// 	if err != nil {
// 		fmt.Printf("Error in creating the file : %v\n", err.Error())
// 		return
// 	}

// 	hash := utils.GenerateTextToHash(encrypt.Password)

// 	encryptedText := utils.Encrypt(fmt.Sprintf("%s", &buffer), hash)
// 	err = utils.WriteToFile(file, encryptedText)
// 	if err != nil {
// 		fmt.Printf("Error in writing the file: %v\n", err.Error())
// 		return
// 	}

// 	fmt.Printf("converted successfully file name is : %v\n", file.Name())
// }
