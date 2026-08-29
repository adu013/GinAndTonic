package main

import (
	"bufio"
	"fmt"
	"gin-and-tonic/models"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	// config.ConnectDatabase()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=============================================")
	fmt.Println("              CREATE A NEW USER              ")
	fmt.Println("=============================================")

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Email Address: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nError reading password")
		return
	}
	password := strings.TrimSpace(string(bytePassword))
	fmt.Println()

	created, err := models.CreateAdminUser(username, email, password)
	if err != nil || !created {
		return
	}

	fmt.Printf("SUCCESS! Created admin user with username: %s", username)

}
