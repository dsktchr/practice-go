package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("期待値: 5, 出力: ", result)
	}
}
