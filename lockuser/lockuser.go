package lockuser

import "fmt"

type Userstore interface {
	username(u adduser.usernmae)
	day(u adduser.day)
}

func lockUser() {
	fmt.Println("Ocserv does not support user locking directly in plain auth mode. Consider deleting user or adding custom logic.")
}
