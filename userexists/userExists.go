package userexists

import (
	"bufio"
	"os"
	"strings"
)

type Userstore interface {
	username(u adduser.username)
}

func userExists(adduser.username, string) (bool, string) {
	file, err := os.Open(userFile)
	if err != nil {
		logger.Println("Error opening user file:", err)
		return false, ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, adduser.username+":") {
			return true, line
		}
	}
	return false, ""
}
