package main

import "fmt"

func outer() {
	var s4 string = "hello"
	fmt.Println(s4)
}

func main() {
	// 変数の勉強
	var i int = 100
	
	fmt.Println(i)

	var s string = "hello"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		a int = 10
    b bool = true
	)
	fmt.Println(a, b)

	// 暗黙的定義
	i4 := 400
	fmt.Println(i4)

	outer()
}