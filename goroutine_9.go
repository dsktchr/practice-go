package main

import (
	"fmt"
	"sync"
)

func processAndGather(processor func(int) int, data []int) []int {
	num := len(data)

	chResult := make(chan int, num)
	// dataと同サイズのバッファを持つチャネルを作成する
	
	var wg sync.WaitGroup
	wg.Add(num)

	for _,v := range data {
		go func(v int) {
			defer wg.Done()
			chResult <- processor(v)
		}(v)
	}
	
	wg.Wait()

	close(chResult)

	var result []int
	// chResultのバッファに残っている要素を処理する
	for v := range chResult {
		result = append(result, v)
	}

	return result
}

func main() {
	inputValues := []int{1,2,3,4,5}
	outputValues := processAndGather(
		func(j int) int { return j*j },
		inputValues)

	fmt.Println(outputValues)
}
