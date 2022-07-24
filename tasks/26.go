package tasks

import "strings"

func Task26() {
	var str1 string = "abcd"
	var str2 string = "abCdefAaf"
	var str3 string = "aabcd"

	PrintResult(26.1, "Уникальные символы", isUnique(str1))
	PrintResult(26.2, "Уникальные символы", isUnique(str2))
	PrintResult(26.3, "Уникальные символы", isUnique(str3))
}

// Бьем строку в массив символов
// затем добавляем символы в мапу
// если в мапе нет одинаковых то добавляем дальше
// если есть совпадение то символы в строке не уникальны
func isUnique(str string) bool {
	var result bool = true
	str = strings.ToLower(str)
	strArr := strings.Split(str, "")

	var uniqueStr map[string]string = map[string]string{}

	for _, val := range strArr {
		if _, ok := uniqueStr[val]; ok {
			result = false
			break
		} else {
			uniqueStr[val] = val
		}
	}

	return result
}
