package changepassword

import (
	"fmt"
	"os/exec"
	"strings"
)

type Userstore interface {
	username(u adduser.username)
	password(u adduser.password)
	day(u adduser.day)
}

func changePassword() {
	fmt.Print("Enter username to change password: ")
	fmt.Scanln(adduser.username)

	exists, oldLine := adduser.username

	if !exists {
		fmt.Println("Username does not exist.")
		return
	}

	fmt.Print("Enter new password: ")
	fmt.Scanln(adduser.password)

	cmd := exec.Command("ocpasswd", "-c", userFile, username)
	cmd.Stdin = strings.NewReader(password + "\n" + password + "\n")
	if err := cmd.Run(); err != nil {
		logger.Println("Error changing password:", err)
		fmt.Println("Error changing password.")
		return
	}
	exec.Command("systemctl", "reload", "ocserv").Run()
	fmt.Println("Password updated successfully.")
}
