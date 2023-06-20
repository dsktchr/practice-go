package table

import "testing"

func TestDoMath(t *testing.T) {
	result, err := DoMath(2, 2, "+")
	if result != 4 {
		t.Error("期待値:4 出力:", result)
	}
	if err != nil {
		t.Error("期待ERROR: nil 出力ERROR:", err)
	}
	result2, err2 := DoMath(2, 2, "-")
	if result2 != 0 {
		t.Error("期待値:0 出力:", result2)
	}
	if err2 != nil {
		t.Error("期待ERROR: nil, 出力ERROR:", err2)
	}
	result3, err3 := DoMath(2, 2, "*")
	if result3 != 4 {
		t.Error("期待値:4 出力:", result3)
	}
	if err3 != nil {
		t.Error("期待ERROR: nil 出力ERROR:", err3)
	}
	result4, err4 := DoMath(2, 2, "/")
	if result4 != 1 {
		t.Error("期待値:1 出力:", result4)
	}
	if err4 != nil {
		t.Error("期待値ERROR: nil 出力ERROR:",err)
	}
}

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name string
		num1 int
		num2 int
		op string
		expected int
		errMsg string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"miltiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, "0除算"},
		{"bad_op", 2, 2, "?", 0, "未知の演算子 ?"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T){
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("期待値: %d 出力: %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("期待ERROR: %s 出力ERROR: %s", d.errMsg, errMsg)
			}
		})

	}
}
