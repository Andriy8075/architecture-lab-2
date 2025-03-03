package lab2

import (
	"testing"
)

func TestEvalPrefix(t *testing.T) {
	tests := []struct {
		expression string
		expected   int
		shouldFail bool
	}{
		{"* 3 4", 15, false},
	}

	for _, test := range tests {
		result, err := EvalPrefix(test.expression)
		if (err != nil) != test.shouldFail {
			t.Errorf("EvalPrefix(%q) несподівана помилка: %v", test.expression, err)
		} else if result != test.expected && !test.shouldFail {
			t.Errorf("EvalPrefix(%q) = %d, очікувалося %d", test.expression, result, test.expected)
		}
		println(result)
	}
}
