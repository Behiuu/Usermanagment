package unlockuser

import (
	"os/exec"
)

const userFile = "/etc/ocserv/ocpasswd"

type Userstore interface {
	username(u adduser.usernmae)
	day(u adduser.day)
}

func unlockUser() {
	cmd := exec.Command("ocpasswd", "-c", userFile, "-u", adduser.username)
	return

}
