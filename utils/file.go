package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName string) (file *os.File, err error) {
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Printf("error in creating the file: %v\n", err)
		return
	}
	return
}

func WriteToFile(file *os.File, content []byte) (err error) {
	if _, err = file.Write(content); err != nil {
		fmt.Printf("error in append the file: %v\n", err)
		return
	}
	return
}

func ReadFromFile(filePath string) (buffer []byte, err error) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error in open the file: %v\n", err)
		return
	}
	defer file.Close()

	// get the state of the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("error in getting the file state: %v\n", err)
		return
	}

	// read the file
	buffer = make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Printf("error in read the file: %v\n", err)
		return
	}

	return
}

func GetFileStat(filePath string) (fileInfo os.FileInfo, err error) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error in open the file: %v\n", err)
		return
	}
	defer file.Close()

	// get the state of the file
	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Printf("error in getting the file state: %v\n", err)
		return
	}

	return
}

func ChangeFileExtension(fileName string, newExtension string) string {
	extension := filepath.Ext(fileName)
	newFileName := fileName[0:len(fileName)-len(extension)] + newExtension
	return newFileName
}
