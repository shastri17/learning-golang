package main

import (
	"fmt"
)

func main() {
	sqFunc := func(a int) int {
		return a * a
	}

	cubeFunc := func(a int) int {
		return a * a * a
	}

	// Following Function takes functions as input
	OutputFunctions([]int{1, 2, 3, 4, 5}, sqFunc, cubeFunc)

	// Similarly, a function can return a function as well.

}

func OutputFunctions(nums []int, functions ...func(int) int) {

	result := make([][]int, len(functions))
	for i, f := range functions {
		result[i] = make([]int, len(nums))
		for j, n := range nums {
			result[i][j] = f(n)
		}
	}

	fmt.Println(result)
}
