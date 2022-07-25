package tasks

import (
	"unicode/utf8"
)

func Task19() {
	var str1 string = "日本語"
	var str2 string = "главрыба"
	var str3 string = "test data"

	PrintInputOutput(19.1, "Переворот строки с помощтю рун (вариант 1)", reverseStrV1(str1), str1)
	PrintInputOutput(19.2, "Переворот строки с помощтю рун (вариант 1)", reverseStrV1(str2), str2)
	PrintInputOutput(19.3, "Переворот строки с помощтю рун (вариант 1)", reverseStrV1(str3), str3)

	PrintInputOutput(19.4, "Переворот строки с помощью пакета utf8 (вариант 2)", reverseStrV2(str1), str1)
	PrintInputOutput(19.5, "Переворот строки с помощью пакета utf8 (вариант 2)", reverseStrV2(str2), str2)
	PrintInputOutput(19.6, "Переворот строки с помощью пакета utf8 (вариант 2)", reverseStrV2(str3), str3)

	PrintInputOutput(19.7, "Переворот строки без учета unicode (вариант 3)", reverseStrV3(str1), str1)
	PrintInputOutput(19.8, "Переворот строки без учета unicode (вариант 3)", reverseStrV3(str2), str2)
	PrintInputOutput(19.9, "Переворот строки без учета unicode (вариант 3)", reverseStrV3(str3), str3)
}

func reverseStrV1(str string) string {
	s := []rune(str)
	var result []rune = make([]rune, 0)

	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return string(result)
}

func reverseStrV2(str string) string {
	var tmp []rune = make([]rune, 0)
	var result []rune = make([]rune, 0)

	for i, w := 0, 0; i < len(str); i += w {
		runeVal, width := utf8.DecodeRuneInString(str[i:])
		w = width
		tmp = append(tmp, runeVal)
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		result = append(result, tmp[i])
	}

	return string(result)
}

func reverseStrV3(str string) string {
	var result string

	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}

	return result
}
