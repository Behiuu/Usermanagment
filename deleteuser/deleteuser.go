package deleteuser

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Userstore interface {
	username(u adduser.usernmae)
}

func deleteUser() {

	fmt.Print("Enter username to delete: ")
	fmt.Scanln(adduser.username)

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
