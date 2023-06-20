package adder_test

import (
	"testing"
	"test_examples/adder3"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("期待値: 5, 出力: ", result)
	}
}
