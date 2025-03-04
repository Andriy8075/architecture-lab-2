package lab2

import (
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	// Читаємо вхідні дані
	data, err := io.ReadAll(h.Input)
	if err != nil {
		return fmt.Errorf("помилка читання: %w", err)
	}

	expression := strings.TrimSpace(string(data))
	if expression == "" {
		return fmt.Errorf("порожній вираз")
	}

	// Викликаємо EvalPrefix (з implementation.go)
	result, err := EvalPrefix(expression)
	if err != nil {
		return fmt.Errorf("помилка обчислення: %w", err)
	}

	// Записуємо результат
	_, err = fmt.Fprintln(h.Output, result)
	return err
}
