package main

import (
	"fmt"
	// "strconv"
	// "os"
	// "time"
)

func main() {
	// var sl []int
	// fmt.Println(sl)

	// var sl2 =  []int{100, 200}
	// fmt.Println(sl2)

	// sl3 := []int{100, 200, 300}
	// fmt.Println(sl3)

	// sl4 := make([]int, 5)
	// fmt.Println(sl4)

	// sl2[0] = 1000
	// fmt.Println(sl2)

	// sl5 := []int{100, 200, 300, 40}
	// fmt.Println(sl5)
	// fmt.Println(sl5[len(sl5)-1])

	// sl := []int{100, 200, 300, 400}
	// fmt.Println(sl)

	// // スライスは要素が可変長
	// sl = append(sl, 500)
	// fmt.Println(sl)

	// sl = append(sl, 600, 700, 800)
	// fmt.Println(sl)

	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// fmt.Println(len(sl2))
	// fmt.Println(cap(sl2))

	// sl3 := make([]int, 5, 10)
	// fmt.Println(sl3)

	// fmt.Println(len(sl3))
	// fmt.Println(cap(sl3))

	// sl3 = append(sl3, 100, 200, 300, 400, 500)
	// fmt.Println(sl3)

	// copy
// 	sl := []int{100, 200, 300, 400, 500}
// 	sl2 := make([]int, 5, 10)
// 	n := copy(sl2, sl)
// // nはコピーした数
// 	fmt.Println(n, sl2)

// sl := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
// for i, v := range sl {
// 	fmt.Println(i, v)

// fmt.Println(Sum(1, 2, 3))

// sl := []int{100, 200, 300, 4}
// fmt.Println(Sum(sl...))

// var m = map[string]int{"A": 100, "B": 200, "C": 300}
// fmt.Println(m)

// m2 := make(map[int]string)
// fmt.Println(m2)

// m2[1] = "japan"
// m2[2] = "korea"
// fmt.Println(m2)
// fmt.Println(m2[1])

// // ハンドリングもできる
// // s, ok := m2[4]
// // if !ok {
// // 	fmt.Println("not found")
// // }
// // fmt.Println(s, ok)

// m2[3] = "china"
// fmt.Println(m2)

// delete(m2, 3)

m := map[string]int{"A": 100, "B": 20}
for i, v := range m {
	fmt.Println(i, v)
}
}

func Sum(s ...int) int {
	total := 0
  for _, v := range s {
    total += v
  }
  return total
}