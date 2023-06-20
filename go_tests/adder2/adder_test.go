package adder

import (
	"testing"
	"time"
	"os"
	"fmt"
)

var testTime time.Time

// TestMainを含んだパッケージを十個すると、先にTestMain関数が実行される
// NOTE: 1パッケージに1つのTestMain関数しか含めることができない
func TestMain(m *testing.M) {
	fmt.Println("テストの準備")
	testTime = time.Now()
	exitVal := m.Run() // ここがテスト関数の呼び出しになる
	fmt.Println("テスト終了")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirstではTestMainで設定されたものを使う", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecondでもestMainで設定されたものを使う", testTime)
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("期待値: 5, 結果: ", result)
	}
}
