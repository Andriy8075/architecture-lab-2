package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalPrefix(t *testing.T) {
	res, err := EvalPrefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func ExampleEvalPrefix() {
	res, _ := EvalPrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
