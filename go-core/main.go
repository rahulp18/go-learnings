package main

import (
	"fmt"
)

type User struct {
	name     string
	email    string
	password string
	age      int
	married  bool
}

func main() {
	// user1 := User{
	// 	name:     "Rahul Pradhan",
	// 	email:    "rahul@gmail.com",
	// 	password: "Rahul",
	// 	age:      23,
	// 	married:  false,
	// }
	var user User
	fmt.Print(user)

}

func changeName(name *string) {
	*name = "Binita's Husband"
}
