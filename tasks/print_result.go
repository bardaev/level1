package tasks

import (
	"fmt"
	"strconv"
)

func PrintResult(name int, result interface{}) {
	fmt.Println("Task " + strconv.Itoa(name))
	fmt.Println(result)

	EndOfTask()
}

func PrintBeforeAfterResult(name int, before interface{}, after interface{}) {
	fmt.Println("Task " + strconv.Itoa(name))

	fmt.Print("Before: ")
	fmt.Println(before)

	fmt.Print("After: ")
	fmt.Println(after)

	EndOfTask()
}

func PrintInputOutput(name float64, inputData interface{}, outputData interface{}) {
	fmt.Print("Task ")
	fmt.Println(name)

	fmt.Print("input data: ")
	fmt.Println(inputData)

	fmt.Print("output data: ")
	fmt.Println(outputData)

	EndOfTask()
}

func EndOfTask() {
	fmt.Println("-------------------------------------")
}
