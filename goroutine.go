package main

import "fmt"

func doBizLogin(val int) int {
	//valを使った処理
	return val + 2
}

func runThingsConcurrectly(chIn <-chan int, chOut chan<- string) {
	for val := range chIn {
		go func(val int) {
			result := doBizLogin(val)
			resultString := fmt.Sprintf("%d -> %d", val, result)
			chOut <- resultString
		}(val)
	}
}

func main() {
	//送信用
	chIn := make(chan int)
	//受信用
	chOut := make(chan string)
	go runThingsConcurrectly(chIn, chOut)

	vals := []int{1,2,3,4,5}
	for _, v := range vals {
		fmt.Println(v)
		chIn <- v
	}

	i := 0
	for v := range chOut {
		fmt.Println(v)
		i++
		if len(vals) <= i {
			break
		}
	}
}
