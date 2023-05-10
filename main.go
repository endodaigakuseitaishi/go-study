package main

import (
	"fmt"
	// "strconv"
	// "os"
)

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
	// x, y float64
}



func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName() {
	u.Name = "tttt"
}

func main() {
	t := T{User: User{Name: "John", Age: 20}}
	t.User.SetName()
	fmt.Println(t)
}
