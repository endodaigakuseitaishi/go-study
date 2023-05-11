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

type MyInt int

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName() {
	u.Name = "tttt"
}

// User型を返すコンストラクタ関数でUser型のポインタを生成
func NewUser(name string, age int) *User {
	return &User{
    Name: name,
    Age:  age,
  }
}

type Users []*User

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)
	fmt.Println(*user1)

	user2 := User{"user2", 20}
	user3 := User{"user3", 30}
	user4 := User{"user4", 40}
	user5 := User{"user5", 50}

	users := Users{}

	users = append(users, &user2, &user3, &user4 ,&user5)
	
	// for _, V := range users {
	// 	fmt.Println(V)
	// }

	users2 := make([]*User, 0)
	users2 = append(users2, &user2, &user3, &user4,&user5)

	for _, V := range users2 {
    fmt.Println(V)
  }

	m := map[int]User{
		1: {Name: "test", Age: 10},
    2: {Name: "test2", Age: 20},
	}
	fmt.Println(m)

	m2 := map[User]string{
		{Name: "test", Age: 10}: "tokyo",
		{Name: "test2", Age: 20}: "chiba",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	m3[1] = User{Name: "test", Age: 10}
	m3[2] = User{Name: "test2", Age: 20}
	m3[3] = User{Name: "test3", Age: 30}
	fmt.Println(m3)

	var mi MyInt
	mi = 10
	fmt.Println(mi)
	fmt.Println(
}
