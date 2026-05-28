package calculate

import "errors"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на 0")
	}
	return float64(a) / float64(b), nil
}

func Multiply(a, b int) int {
	return a * b
}
