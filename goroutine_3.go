package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func() {
		v := 1
		fmt.Println("この下でch1へ%vを入れる", v)
		ch1 <- v
	}()

	go func() {
		v := 200
		fmt.Println("この下でch2へ%vを入れる", v)
		ch2 <- v
	}()

	go func() {
		//ch3から値を受け取る
		v := <-ch3
		fmt.Println("ch3から値を受け取った", v)
	}()

	go func() {
		v := 4
		ch4 <- v
	}()

	x := 34
	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2", v)
	case ch3 <- x:
		fmt.Println("ch3へ書き込み", x)
	case <-ch4:
		fmt.Println("ch4から値を貰ったが、値を無視する")
	}
}
