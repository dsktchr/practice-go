package main

import "fmt"


func main() {

	// 配列がある
	a := []int{2,4,6,8,10,12,14,16,18,20}
	ch := make(chan int, len(a))

	// 配列の値をチャネルへ送信する
	for _, v := range a {
		//v := v //シャドーイングして…
		go func(val int) {
			//ch <- v * 3
			ch <- val * 2 // 引数で渡されるパターン
		}(v)
	}

	// 配列の長さ分ループする
	for i := 0; i < len(a); i++ {
		// チャネルから値を取得する
		fmt.Print(<-ch, " ")
	}

	fmt.Println()
}
