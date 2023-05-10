package main

import (
	"fmt"
	// "strconv"
	// "os"
)

func main() {
	var n int = 100
	// fmt.Println(n)

	// fmt.Println(&n)

	// Double(n)
	// fmt.Println(n)

	var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// *p = 300
	// n = 200
  Doublev2(&n)
	fmt.Println(n)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3, 4, 5, 6}
	Doublev3(sl)
	fmt.Println(sl)
}

func Doublev3(s []int) {
  for i, v := range s {
		s[i] = v * 2
	}
}

func Double(i int) {
	i = i*2
}

func Doublev2(i *int) {
	*i = *i * 2
}


