package main

import (
	"fmt"
	// "strconv"
	// "os"
)

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
  return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type myError struct {
	Message string
	ErrorCode int
}

type Point struct {
	A int
	B string
}

func(e *myError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &myError{
    Message: "Something went wrong",
    ErrorCode: 100,
  }
}

func main() {
	vs := []Stringfy{
		&Person{
      Name: "John Doe",
      Age:  18,
    },
    &Car{
      Number: "123456789",
      Model:  "Brown",
    },
	}

	for _, v := range vs {
    fmt.Println(v.ToString())
  }

	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*myError)
	if ok {
    fmt.Println(e.ErrorCode)
  }

	p := &Point{100, "123"}
	fmt.Println(p)
}