# go-study
##  channel references for zenn
https://zenn.dev/hsaki/books/golang-concurrency/viewer/chaninternal

## address references for qiita
https://qiita.com/gold-kou/items/0d1585fb8d687061c890

## channel 
### ch定義
```
ch10 := make(chan int)
ch11 := make(chan int)
ch12 := make(chan int)
```

### reciever記述
```
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
```
無名関数のゴルーチンで並行処理

### mainの処理部分
```
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
```
