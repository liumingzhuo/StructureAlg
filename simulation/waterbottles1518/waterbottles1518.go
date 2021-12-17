package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	r := numBottles
	for numBottles >= numExchange {
		e := numBottles / numExchange
		y := numBottles % numExchange
		numBottles = e + y
		r += e
	}
	return r
}
func main() {
	ans := numWaterBottles(15, 4)
	fmt.Println(ans)
}
