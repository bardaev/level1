package tasks

import (
	"fmt"
	"strconv"
)

func PrintResult(name float64, description string, result interface{}) {
	fmt.Print("Task ")
	fmt.Println(name)
	fmt.Println("Description: " + description)
	fmt.Print("output data: ")
	fmt.Println(result)

	endOfTask()
}

func PrintBeforeAfterResult(name int, description string, before interface{}, after interface{}) {
	fmt.Println("Task " + strconv.Itoa(name))
	fmt.Println("Description: " + description)

	fmt.Print("Before: ")
	fmt.Println(before)

	fmt.Print("After: ")
	fmt.Println(after)

	endOfTask()
}

func PrintInputOutput(name float64, description string, outputData interface{}, inputData ...interface{}) {
	fmt.Print("Task ")
	fmt.Println(name)

	fmt.Println("Description: " + description)

	for idx, val := range inputData {
		fmt.Print("input data " + strconv.Itoa(idx+1) + ": ")
		fmt.Println(val)
	}

	fmt.Print("output data: ")
	fmt.Println(outputData)

	endOfTask()
}

func PrintConcurrent(name float64, description string, cb func()) {
	fmt.Print("Task ")
	fmt.Println(name)

	fmt.Println("Description: " + description)

	cb()

	endOfTask()
}

func endOfTask() {
	fmt.Println("-------------------------------------")
}
