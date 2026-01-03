package main

import (
	"fmt"

	"github.com/Swapnilden/podcast/auth"
	"github.com/Swapnilden/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("Swapnilden", "Secret")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@example.com",
		// Name:  "John Doe",
	}

	color.Green(user.Email)
}
