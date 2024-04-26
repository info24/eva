package serve

import (
	"github.com/howeyc/gopass"
	"github.com/info24/eva/service"
	"os"
)

func AddUser(username string) error {
	password, err := gopass.GetPasswdPrompt("password: ", true, os.Stdin, os.Stdout)
	if err != nil {
		println("Error reading from stdin")
		println(err.Error())
		os.Exit(-1)
	}

	err = service.CreateUser(username, string(password))
	if err != nil {
		println("Error creating user")
		println(err.Error())
		os.Exit(-1)
	}
	println("user " + username + " created.")
	return nil
}
