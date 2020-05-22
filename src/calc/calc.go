package calc

import "errors"

func Sum(a, b int) int {
	return a + b
}

func Div(a, b int) (int, error) {

	if b == 0 {
		return -1, errors.New("can't work with 0")
	}

	return a / b, nil
}

func Res(a, b int) int {
	return a - b
}

func Multi(a, b int) int {
	return a * b
}
