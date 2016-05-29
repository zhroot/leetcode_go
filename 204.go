/**
204. Count Primes
**/
package main

import (
	"fmt"
	"math"
)

func InvSqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	return 1 / x
}

//筛选法
func countPrimes(n int) int {
	n = n - 1
	slice := make([]int, n+1, n+1)
	var curIndex = 2
	var maxIndex = int(InvSqrt(float32(n)))
	var result = 0
	var i int
	for {
		if curIndex > maxIndex {
			break
		}
		if slice[curIndex] != 0 {
			curIndex = curIndex + 1
			continue
		}
		for i = 2; i <= n/curIndex; i++ {
			slice[i*curIndex] = 1
		}
		curIndex = curIndex + 1
	}

	for i = 2; i < n+1; i++ {
		if slice[i] == 0 {
			result = result + 1
		}
	}
	return result

}

func main() {
	fmt.Println(countPrimes(100000))
}
