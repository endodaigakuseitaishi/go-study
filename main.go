package main

import (
	"fmt"
	// "strconv"
	// "os"
	"time"
)

func main() {
	// a := 0
	// if a == 2 {
	// 	fmt.Println("a is 2")
	// } else if a == 1 {
	// 	fmt.Println("a is not 1")
	// } else {
	// 	fmt.Println("a is neither 1 nor 2")
	// }

	// if b := 100; b == 100 {
	// 	fmt.Println("b is 100")
  // }

	// エラーハンドリング
	// var s string = "error"

	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// fmt.Printf("i = %T\n", i)

	// for
	// i := 0
	// for {
	// 	i++
	// 	if i == 10 {
  //     break
  //   }
	// 	fmt.Println(i+1)
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
  //   point++
	// }

	// arr := [3]int{1, 2, 3}
	// for i:=0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1, 2, 3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"a", "b", "c"}
	// for i, v := range sl {
  //   fmt.Println(i, v)
  // }

	// m := map[string]int{"a": 1, "b": 2, "c" : 3}
	// for k, v := range m {
  //   fmt.Println(k, v)
  // }

	// var x interface{} = 3.41
	// i := x.(int)
	// fmt.Println(i + 2)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f + 2, isFloat64)

	// if x == nil {
	// 	fmt.Println("none")
	// } else if i, isInt := x.(int); isInt {
	// 	fmt.Println(i + 2)
	// } else if f, isFloat := x.(float64); isFloat {
	// 	fmt.Println(f + 2, isFloat)
  // } else if s, isString := x.(string); isString {
	// 	fmt.Println(s)
	// } else {
	// 	fmt.Println("unknown")
	// }

	// anything(3.11)

	// defer
	// TestDefer()

	// // defer func() {
	// // 	fmt.Println("1")
	// // 	fmt.Println("2")
	// // 	fmt.Println("3")
	// // }()

	// RunDefer()

	// file, err := os.Create("test.go")
	// if err!= nil {
  //   panic(err)
  // }
	// defer file.Close()
	// file.Write([]byte("package main\n"))

	// goroutin
	go sub()
	// go sub()
	for {
		fmt.Println("main")
		time.Sleep(200 * time.Millisecond)
	}
}

// 初期化処理
// func init() {
// 	fmt.Println("init")
// }
// func init() {
// 	fmt.Println("reinit")
// }

func sub() {
	for {
		fmt.Println("sub")
		time.Sleep(100 * time.Millisecond)
	}
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func TestDefer() {
	defer fmt.Println("defer")
	fmt.Println("START")
}

func anything(a interface{}) {
	switch v:= a.(type) {
	case int:
		fmt.Println(v)
	case float64:
		fmt.Println(v + 1222)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
}
}