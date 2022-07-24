package tasks

import "strconv"

func Task13() {
	var a int = 2
	var b int = 5

	resA, resB := replace(a, b)

	PrintInputOutput(
		13.0,
		"a = "+strconv.Itoa(a)+", b = "+strconv.Itoa(b),
		"a = "+strconv.Itoa(resA)+", b = "+strconv.Itoa(resB))
}

func replace(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
