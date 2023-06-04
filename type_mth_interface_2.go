package main

import (
	"fmt"
	"time"
)

func main() {
	// ゼロ値で宣言
	var c Counter

	fmt.Println(c.String())

	c.Increment()

	fmt.Println(c.String())
}

type Counter struct {
	total int
	lastUpdated time.Time
}

// ポインタレシーバー（レシーバーの値を変更するため）
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// 値レシーバー（レシーバーの値を変更しないため）
func (c Counter) String() string {
	return fmt.Sprintf("合計=%d, 更新時刻: %v", c.total, c.lastUpdated)
}
