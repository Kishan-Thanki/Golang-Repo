package main

import (
	"fmt"

	"github.com/Kishan-Thanki/Golang-Repo/auth"
	"github.com/Kishan-Thanki/Golang-Repo/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("LionelMessi", "secretpassword")

	session := auth.GetSession()
	fmt.Println("Session:", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}
	// fmt.Println(user.Email, user.Name)
	color.Red(user.Email)
	color.Blue(user.Name)
}
