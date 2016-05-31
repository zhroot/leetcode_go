/*
*134. Gas Station
 */
package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	var length = len(gas)
	var i int
	for i = 0; i < length; i++ {
		var start1 = i
		var end1 = length
		var start2 = 0
		var end2 = i
		var j int
		var remain = 0
		for j = start1; j < end1; j++ {
			remain = remain + gas[j] - cost[j]
			if remain < 0 {
				break
			}
		}
		if remain < 0 {
			continue
		}
		for j = start2; j < end2; j++ {
			remain = remain + gas[j] - cost[j]
			if remain < 0 {
				break
			}
		}

		if remain >= 0 {
			return i
		}
	}
	return -1
}

/*
* test
 */

func makecase(n int) ([]int, []int) {
	var gas = []int{1, 2, 3, 4, 5, 6}
	var cost = []int{6, 5, 4, 3, 2, 1}
	return gas, cost
}
func main() {
	var gas []int
	var cost []int
	gas, cost = makecase(10)
	var result = canCompleteCircuit(gas, cost)
	fmt.Println(result)
}
