package main

import (
	"fmt"
)

type User struct {
	Name string
}

func Scan(u *User, name string) {
	u.Name = name
}

func SetNameByValue(name string) User {
	return User{Name: name}
}

func main() {
	var u User
	fmt.Println(u)
	Scan(&u, "hai")
	fmt.Println(u)
	u2 := SetNameByValue("hai2")
	fmt.Println(u2)
}
