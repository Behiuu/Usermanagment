package adduser

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type adduser struct {
	username string
	password string
	day      int
}

func addUser(username, password string) string {
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

	expiration := time.Now().Add(time.Duration(days) * 24 * time.Hour).Format("2006-01-02")

	cmd := exec.Command("ocpasswd", "-c", userFile, username)
	cmd.Stdin = strings.NewReader(password + "\n" + password + "\n")
	if err := cmd.Run(); err != nil {
		logger.Println("Error adding user with ocpasswd:", err)
		fmt.Println("Error creating user.")
		return
	}

	appendExpiration(username, expiration)
	exec.Command("systemctl", "reload", "ocserv").Run()
	fmt.Println("User added successfully.")
}
