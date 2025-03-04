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
		{"* 1 bebra", 0, true}, // Некоректні символи
	}

	for _, test := range tests {
		result, err := EvalPrefix(test.expression)

		println(result, err)

		if (err != nil) != test.shouldFail {
			t.Errorf("EvalPrefix(%q) несподівана помилка: %v", test.expression, err)
		} else if result != test.expected && !test.shouldFail {
			t.Errorf("EvalPrefix(%q) = %d, очікувалося %d", test.expression, result, test.expected)
		}
		//println(result)
	}
}
