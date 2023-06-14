package main

import "fmt"

type users struct {
	name     string
	email    string
	password string
}

var sliceUser = []users{}

func register() {
	usersX := users{
		name:     "John Doe",
		email:    "John@gmail.com",
		password: "abcde",
	}

	sliceUser = append(sliceUser, usersX)

	usersX2 := users{
		name:     "John Doe 2",
		email:    "John@gmail.com",
		password: "abcde",
	}

	sliceUser = append(sliceUser, usersX2)

	usersX3 := users{
		name:     "John Doe 3",
		email:    "John@gmail.com",
		password: "abcde",
	}

	sliceUser = append(sliceUser, usersX3)
}

func get() {
	fmt.Println(sliceUser)
}
func main() {
	register()
	get()
}
