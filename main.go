package main

import (
	"fmt"
	// "strconv"
	// "os"
)

func main() {
	// channel
	// var ch1 chan int 

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch2 chan<- int

	// ch1 = make(chan int)
	// ch2 := make(chan int)

	// fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	// ch3 := make(chan int, 5)
	// fmt.Println(cap(ch3))

	// ch3に送信
	// ch3 <- 1
	// fmt.Println(len(ch3))

	// iに送信
	// i := <-ch3
	// fmt.Println(i)
	// i2 := <-ch3
	// fmt.Println(i2)

	// ch3 <- 1
	// fmt.Println("========")
	// fmt.Println(len(ch3))
	// fmt.Println(<-ch3)
	// ch3 <- 2
	// ch3 <- 3
	// ch3 <- 4
	// ch3 <- 5
	// ch3 <- 6

	// ch4 := make(chan int)
	// ch5 := make(chan int)

	// go reciever(ch4)
	// go reciever(ch5)

	// i := 0

	// for i < 100 {
	// 	ch4 <- i
	// 	ch5 <- i
	// 	time.Sleep(50*time.Millisecond)
	// 	i++
	// }

	// ch6 := make(chan int, 2)
	// go reciever("1.",ch6)
	// go reciever("2.",ch6)
	// go reciever("3.",ch6)

	// i := 0
	// for i < 100 {
  //   ch6 <- i
  //   i++
  // }
	// close(ch6)
	// time.Sleep(3*time.Second)

	// chが空かつchがcloseでfalseを返す
	// i, ok := <-ch6
	// fmt.Println(i, ok)

	// ch7 := make(chan int, 3)
	// ch7 <- 1
	// ch7 <- 2
	// ch7 <- 3
  // close(ch7)

	// for i := range ch7 {
	// 	fmt.Println(i)
	// }

	// ch8 := make(chan int, 2)
	// ch9 := make(chan string, 2)

	// ch9 <- "a"
	// ch9 <- "b"
	// ch8 <- 2000
	// ch8 <- 1000
	// v1 := <-ch8
	// v2 := <-ch9
	// fmt.Println(v1)
	// fmt.Println(v2)

	// selectの中はランダム選択
	// select {
	// case v1 := <-ch8:
	// 	fmt.Println(v1 + 1000)
	// case v2 := <-ch9:
	// 	fmt.Println(v2 + "AAA")
	// default:
	// 	fmt.Println("どちらでも")
	// }

	ch10 := make(chan int)
	ch11 := make(chan int)
	ch12 := make(chan int)

	// reciever
	go func() {
		for{
			i := <-ch10
			ch11 <- i*2
		}
	}()

	go func() {
		for{
			i2 := <-ch11
			ch12 <- i2-2
		}
	}()
	
	n := 0
	for {
		select {
			case ch10 <- n:
				n++
			case i3 := <-ch12:
				fmt.Println("received!!!", i3)
		}
		if n >= 100 {
      break
    }
	}
}

// func reciever(c chan int) {
// 	for {
// 		i := <-c
// 		fmt.Println(i + 3)
// 	}
// }

// func reciever(name string, c chan int) {
// 	for {
// 		i, ok := <-c
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(name, i)
// 	}
// 	fmt.Println(name, "done")
// }


