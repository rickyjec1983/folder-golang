package main

import "fmt"

type users struct {
	name     string
	email    string
	password string
}

var sliceUser = []users{}

func register(namex string, emailx string, passwordx string) {
	usersX := users{
		name:     namex,
		email:    emailx,
		password: passwordx,
	}

	sliceUser = append(sliceUser, usersX)

}

func get() {
	fmt.Println(sliceUser)
}
func main() {
	register("ricky 1", "ricky@gmail.com", "q123")
	register("ricky 2", "ricky@gmail.com", "q123")
	register("ricky 3", "ricky@gmail.com", "q123")

	get()
}
