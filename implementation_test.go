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
		{"^ 1 2", 0, true},
		{"- 10 5", 5, false},
		{"* 3 4", 12, false},
		{"* 2 a", 0, true}, // Некоректні символи
		{"/ 5 0", 0, true}, // Ділення на нуль
		{"- + + + * 1 4 3 3 / 16 - 7 5 20", -2, false},
		{"^ + * - 20 5 / 16 4 + 8 - 6 2 2", 5184, false},
		{"- / + 30 10 * 3 4 - 18 6", -9, false},
		{"* - + 25 5 10 / 20 2 + 3 1", 30, true},
		{"+ + * 6 3 - 12 4 / 24 6", 26, true},
		{"- * + 7 5 3 / 32 8 + 10 2", 36, true},
	}

	for _, test := range tests {
		result, err := EvalPrefix(test.expression)

		println(result, err)

		if err != nil && !test.shouldFail {
			t.Errorf("EvalPrefix(%q) несподівана помилка: %v", test.expression, err)
		} else if result != test.expected && !test.shouldFail {
			t.Errorf("EvalPrefix(%q) = %d, очікувалося %d", test.expression, result, test.expected)
		} else if result == test.expected && test.shouldFail && err == nil {
			t.Errorf("Очікувалася помилка, якої немає")
		}
		//println(result)
	}
}
