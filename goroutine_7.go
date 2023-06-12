package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
)

func main() {
	fmt.Println(">> main開始")
	rand.Seed(time.Now().UnixNano())
	funcs := prepareFunctions() // 関数のスライスを作成する
	s := "Time files like an arrow"
	fmt.Println("オリジナル", s)

	r := convertDate(s, funcs)
	fmt.Printf("一番早く到達した結果", r)

	fmt.Println("他のごルーチンの完了を待つ")
	time.Sleep(1*time.Second)
	fmt.Println("<< main終了")

}

type message struct {
	s string
	fromFunc int
}

func convertDate(s string, converters []func(string) message) message {
	done := make(chan struct{})
	resultChan := make(chan message)
	for _, f := range converters {
		go func(f func(string) message) {
			r := f(s)
			select {
			case resultChan <- r:
				fmt.Printf("結果が戻ってきたので %v\n", r)
			case <- done:
				fmt.Println("case <-done選択", r.fromFunc)
			}
		}(f)
	}

	r := <-resultChan
	close(done)
	return r
}

func randomPeriod() time.Duration {
	t := time.Millisecond * time.Duration(rand.Intn(4)+2)
	return t
}

func prepareFunctions() []func(string) message {
	funcNo1 := func(a string) message {
		fmt.Println(">> funcNo1開始")
		sleepRandomPeriod(1)
		b := strings.ToLower(a)
		fmt.Println("<< funcNo1終了")
		return message{b, 1}
	}

	funcNo2 := func(a string) message {
		fmt.Println(">> funcNo2開始")
		sleepRandomPeriod(2)
		b := strings.ToUpper(a)
		fmt.Println("<< funcNo2終了")
		return message{b, 2}
	}

	funcNo3 := func(a string) message {
		fmt.Println(">> funcNo3開始")
		sleepRandomPeriod(3)
		b := strings.ReplaceAll(a, "i", "I")
		fmt.Println("<< funcNo3終了")
		return message{b, 3}
	}

	funcNo4 := func(a string) message {
		fmt.Println(">> funcNo4開始")
		sleepRandomPeriod(4)
		b := strings.ReplaceAll(a, "e", "E")
		a = strings.ReplaceAll(a, "r", "L")
		fmt.Println("<< funcNo4終了")
		return message{b, 4}
	}

	return []func(string) message{funcNo1,funcNo2,funcNo3,funcNo4}
}

func sleepRandomPeriod(funcNum int) {
	t := randomPeriod()
	fmt.Println(funcNum, "will sleep: ", t)
	time.Sleep(t)	
}
