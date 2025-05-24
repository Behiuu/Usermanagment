package main

import (
	"fmt"
	"log"
	"os"
)

const userFile = "/etc/ocserv/ocpasswd"
const logFile = "ocservcli.log"

var logger *log.Logger

func init() {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	fmt.Println("Welcome To Ocserv Username Managment")
	fmt.Println("v1.1.0")
	for {
		fmt.Println("Choose an action:")
		fmt.Println("1) Add User")
		fmt.Println("2) Extend User")
		fmt.Println("3) Delete User")
		fmt.Println("4) Change Password")
		fmt.Println("5) Lock User")
		fmt.Println("6) Unlock User")
		fmt.Println("0) Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addUser()
		case 2:
			extendUser()
		case 3:
			deleteUser()
		case 4:
			changePassword()
		case 5:
			lockUser()
		case 6:
			unlockUser()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
