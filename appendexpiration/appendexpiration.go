package appendexpiration

import (
	"bufio"
	"os"
	"strings"
)

func appendExpiration(username, expiration string) {
	file, err := os.OpenFile(userFile, os.O_RDWR, 0600)
	if err != nil {
		logger.Println("Error opening user file for expiration:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, username+":") {
			line = line + ":" + expiration
		}
		content = append(content, line)
	}

	file.Seek(0, 0)
	file.Truncate(0)
	for _, l := range content {
		file.WriteString(l + "\n")
	}
}
