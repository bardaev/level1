package tasks

func Task14() {
	var num int = 5
	var str string = "str"
	var boolean bool = true
	var ch chan interface{} = make(chan interface{})

	PrintInputOutput(14.1, "Определение типа в рантайме", getType(num), num)
	PrintInputOutput(14.2, "Определение типа в рантайме", getType(str), str)
	PrintInputOutput(14.3, "Определение типа в рантайме", getType(boolean), boolean)
	PrintInputOutput(14.4, "Определение типа в рантайме", getType(ch), ch)
}

func getType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	default:
		return "unknown"
	}
}
