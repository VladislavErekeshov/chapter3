package main

import "fmt"

func main() {
	x := [5]float64{
		57,
		56,
		58,
		55,
		59,
	}

	var ans float64
	ans = 0
	for _, value := range x {
		ans += value
	}
	fmt.Println(ans / float64(len(x)))
}
