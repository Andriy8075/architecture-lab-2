package lab2

import (
	"errors"
	"strconv"
	"strings"
)

func EvalPrefix(expression string) (int, error) {
	tokens := strings.Fields(expression)
	if len(tokens) == 0 {
		return 0, errors.New("порожній вираз")
	}

	var stack []string
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if isOperator(token) {
			if len(stack) < 2 {
				return 0, errors.New("некоректний вираз")
			}
			op1, _ := strconv.Atoi(stack[len(stack)-1])
			op2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]

			result, err := applyOperator(token, op1, op2)
			if err != nil {
				return 0, err
			}

			stack = append(stack, strconv.Itoa(result))
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("некоректний вираз")
	}

	return strconv.Atoi(stack[0])
}

func applyOperator(operator string, a, b int) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("ділення на нуль")
		}
		return a / b, nil
	case "^":
		result := 1
		for i := 0; i < b; i++ {
			result *= a
		}
		return result, nil
	default:
		return 0, errors.New("невідомий оператор")
	}
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}
