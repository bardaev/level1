package tasks

// Без комментариев
func Task12() {
	var arr []string = []string{"cat", "cat", "dog", "cat", "tree"}
	var result map[string]string = make(map[string]string)

	for _, val := range arr {
		result[val] = val
	}

	PrintInputOutput(12, "Map из slice", result, arr)
}
