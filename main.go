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

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "John Doe"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	user3 := User{Name: "aaa", Age: 12}
	fmt.Println(user3)

	// 構造体の順番に従う必要あり
	// User{12, "test"}はNG
	user4 := User{"bbb", 13}
	fmt.Println(user4)

	user5 := User{Name: "ccc"}
	fmt.Println(user5)

	// 構造体のポインタ型を返す
	user6 := new(User)
	fmt.Println(user6)

	user7 := &User{}
	fmt.Println(user7)

	UpdateUser(user1)
	UpdateUser2(user7)

	fmt.Println(user1)
	fmt.Println(user7)
}

// 構造体
func UpdateUser(user User) {
	user.Name = "test"
	user.Age = 12
}

// 参照渡し
func UpdateUser2(user *User) {
	user.Name = "test"
	user.Age = 12
}