package main

import "fmt"

func main() {
	// 算術演算子
	fmt.Println(1 + 3)
	fmt.Println(1 - 3)
	fmt.Println(1 * 3)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)

	// 比較演算子
	fmt.Println(1 == 1)
	fmt.Println(1!= 1)
	fmt.Println(3 > 1)

	// 論理演算子
	fmt.Println(true || true == false)
}