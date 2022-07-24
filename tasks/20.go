package tasks

import "strings"

func Task20() {
	var str string = "snow dog sun"

	var result string = reverseWord(str)

	PrintInputOutput(20.0, "", result, str)
}

// Чтобы разместить слова в обратном порядке
// можно разбить строку на массив
// и соеденить обратно в строку в обратном порядке
func reverseWord(words string) string {
	var splitWords []string = strings.Split(words, " ")
	var result []string = make([]string, len(splitWords))

	var j int = 0

	for i := len(splitWords) - 1; i >= 0; i-- {
		result[j] = splitWords[i]
		j++
	}

	return strings.Join(result, " ")
}
