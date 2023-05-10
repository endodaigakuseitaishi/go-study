package main

import (
	"fmt"
	// "strconv"
	// "os"
)

type User struct {
	Name string
	Age  int
	// x, y float64
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "test"}
	user1.SayName()

	user1.SetName("aaaaaaaaaaaaaaaaaaaa")
	user1.SayName()

	user2 := &User{Name: "test333"}
	user2.SetName("aaaaaaaaaaaa")
	user2.SayName()
}
