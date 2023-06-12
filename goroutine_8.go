package main

import (
	"fmt"
	"sync"
)

func main() {
	// このwgを参照するようにする
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// 👆のwgをキャプチャする
		defer wg.Done()
		doThing1()
	}()


	// 関数に値のコピーを渡してしまうと、カウントできない
	//go func(wg sync.WaitGroup) {
	//	defer wg.Done()
	//	doThing1()
	//}(wg)

	go func() {
		defer wg.Done()
		doThing2()
	}()

	go func() {
		defer wg.Done()
		doThing3()
	}()

	// Done で3回デクリメントされたらポーズが終了する
	wg.Wait()
}

func doThing1() {
	fmt.Println("Thing 1 done")
}

func doThing2() {
	fmt.Println("Thing 2 done")
}

func doThing3() {
	fmt.Println("Thing 3 done")
}
