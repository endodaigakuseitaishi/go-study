package main

import "fmt"

func main() {
	// 変数の勉強
	// 不動小数点
	var fl64 float64 = 3.14
	fmt.Println(fl64)

	// bool
	var b bool = true
  fmt.Println(b)

	// 配列
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	var arr3 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Println(arr3[2])

	// 配列要素数カウント
	fmt.Println(len(arr1))
}