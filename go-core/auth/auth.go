package auth

import "fmt"

func RegisterUser(email string, password string) {
	fmt.Println("User with " + email + " and password is " + password + " is added")
}
