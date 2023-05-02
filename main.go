package main

import (
	"fmt"
)

// func plus(x int, y int) int {
// 	return x + y
// }

// func div(x, y int) (int, int) {
// 	q := x / y
// 	r := x % y
// 	return q, r
// }

// func double(price int) (result int) {
// 	// resultは返り値
// 	result = price * 2
//   return
// }

// 関数を返す関数
// func returnFunc() func()  {
// 	return func() {
// 		fmt.Println("return func()")
// 	}
// }

// クロージャー
// func later() func(string) string {
// 	var store string
// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s
// 	}
// }

// ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
    i++
    return i
  }
}

func main() {
	// i := plus(1, 2)
	// fmt.Println(i)

	// i2, i3 := div(1, 2)
	// fmt.Println(i2, i3)

	// i4 := double(100)
	// fmt.Println(i4)

	// 無名関数
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// i := f(1, 2)
	// fmt.Println(i)

	// f := returnFunc()
	// f()

	// f := later()
	// fmt.Println(f("a"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}