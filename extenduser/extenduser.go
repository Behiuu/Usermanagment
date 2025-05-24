package extenduser

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func extendUser() {
	var username string
	var extraDays int
	fmt.Print("Enter username to extend: ")
	fmt.Scanln(&username)

	exists, oldLine := userExists(username)
	if !exists {
		fmt.Println("Username does not exist.")
		return
	}

	fmt.Print("How many extra days to extend? ")
	fmt.Scanln(&extraDays)

	parts := strings.Split(oldLine, ":")
	if len(parts) < 4 {
		fmt.Println("User entry missing expiration. Cannot extend.")
		return
	}

	exp, err := time.Parse("2006-01-02", parts[3])
	if err != nil {
		fmt.Println("Invalid expiration format.")
		return
	}

	newExp := exp.Add(time.Duration(extraDays) * 24 * time.Hour).Format("2006-01-02")

	fileContent, err := os.ReadFile(userFile)
	if err != nil {
		logger.Println("Error reading user file:", err)
		fmt.Println("Error reading user file.")
		return
	}

	lines := strings.Split(string(fileContent), "\n")
	updated := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line, username+":") {
			parts := strings.Split(line, ":")
			if len(parts) >= 4 {
				parts[3] = newExp
				line = strings.Join(parts, ":")
			}
		}
		if strings.TrimSpace(line) != "" {
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
	fmt.Println("User expiration extended successfully.")
}
