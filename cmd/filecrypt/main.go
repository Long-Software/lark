package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Long-Software/lark/cmd/filecrypt/encrypt"
	"golang.org/x/term"
)

// This is a file encryption and decryption services
// params:
// 		Type: Encrypt|Decrypt
// 		Path: path_to_file

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]
	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	}

}

func printHelp() {}
func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the file path")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}

	password := getPassword()
	fmt.Println("encrypting")
	err := encrypt.Encrypt(file, password)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("file successfully encrypt")
}
func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the file path")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}

	password := getPassword()
	fmt.Println("decrypting")
	err := encrypt.Decrypt(file, password)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("file successfully encrypt")
}

func getPassword() []byte {
	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Println("Confirm password: ")
	confirm, _ := term.ReadPassword(1)
	if !validatePassword(password, confirm) {
		fmt.Println("Passwords do not match")
		return getPassword()
	}
	return password
}

func validatePassword(password, confirm []byte) bool {
	return bytes.Equal(password, confirm)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
