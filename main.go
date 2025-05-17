package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

func userExists(username string) (bool, string) {
	file, err := os.Open(userFile)
	if err != nil {
		logger.Println("Error opening user file:", err)
		return false, ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, username+":") {
			return true, line
		}
	}
	return false, ""
}

func addUser() {
	var username, password string
	var days int
	fmt.Print("Enter new username: ")
	fmt.Scanln(&username)

	exists, _ := userExists(username)
	if exists {
		fmt.Println("Username already exists.")
		return
	}

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	fmt.Print("How many days is the account valid for? ")
	fmt.Scanln(&days)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Println("Error generating password hash:", err)
		fmt.Println("Internal error.")
		return
	}

	entry := fmt.Sprintf("%s:%s\n", username, string(hash))

	file, err := os.OpenFile(userFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		logger.Println("Error opening user file:", err)
		fmt.Println("Error opening user file.")
		return
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	if err != nil {
		logger.Println("Error writing to user file:", err)
		fmt.Println("Error writing to user file.")
		return
	}

	exec.Command("systemctl", "reload", "ocserv").Run()
	fmt.Println("User added successfully.")
}

func extendUser() {
	fmt.Println("Note: Expiration management must be enforced manually or through external logic.")
}

func deleteUser() {
	var username string
	fmt.Print("Enter username to delete: ")
	fmt.Scanln(&username)

	exists, _ := userExists(username)
	if !exists {
		fmt.Println("Username does not exist.")
		return
	}

	file, err := os.ReadFile(userFile)
	if err != nil {
		logger.Println("Error reading user file:", err)
		fmt.Println("Error reading user file.")
		return
	}

	lines := strings.Split(string(file), "\n")
	filtered := []string{}
	for _, line := range lines {
		if !strings.HasPrefix(line, username+":") && strings.TrimSpace(line) != "" {
			filtered = append(filtered, line)
		}
	}

	err = os.WriteFile(userFile, []byte(strings.Join(filtered, "\n")), 0600)
	if err != nil {
		logger.Println("Error writing user file:", err)
		fmt.Println("Error updating user file.")
		return
	}

	exec.Command("systemctl", "reload", "ocserv").Run()
	fmt.Println("User deleted successfully.")
}

func changePassword() {
	var username, password string
	fmt.Print("Enter username to change password: ")
	fmt.Scanln(&username)

	exists, _ := userExists(username)
	if !exists {
		fmt.Println("Username does not exist.")
		return
	}

	fmt.Print("Enter new password: ")
	fmt.Scanln(&password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Println("Error generating password hash:", err)
		fmt.Println("Internal error.")
		return
	}

	file, err := os.ReadFile(userFile)
	if err != nil {
		logger.Println("Error reading user file:", err)
		fmt.Println("Error reading user file.")
		return
	}

	lines := strings.Split(string(file), "\n")
	updated := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line, username+":") {
			updated = append(updated, fmt.Sprintf("%s:%s", username, string(hash)))
		} else if strings.TrimSpace(line) != "" {
			updated = append(updated, line)
		}
	}

	err = os.WriteFile(userFile, []byte(strings.Join(updated, "\n")), 0600)
	if err != nil {
		logger.Println("Error writing user file:", err)
		fmt.Println("Error updating user file.")
		return
	}

	exec.Command("systemctl", "reload", "ocserv").Run()
	fmt.Println("Password updated successfully.")
}

func lockUser() {
	fmt.Println("Ocserv does not support user locking directly in plain auth mode. Consider deleting user or adding custom logic.")
}

func unlockUser() {
	fmt.Println("Ocserv does not support unlocking directly in plain auth mode.")
}
