/*
* 118. Pascal's Triangle
 */
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	var result [][]int

	var i, j int

	for i = 1; i <= numRows; i++ {
		var tempResult []int
		tempResult = append(tempResult, 1)
		for j = 1; j < i-1; j++ {
			tempResult = append(tempResult, result[i-2][j-1]+result[i-2][j])
		}
		if i != 1 {
			tempResult = append(tempResult, 1)
		}
		result = append(result, tempResult)
	}
	return result
}

func main() {
	var result [][]int = generate(10)
	fmt.Println(result)
}
