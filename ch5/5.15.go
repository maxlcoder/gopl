package main

import "fmt"

func max(max int, vals ...int) int {
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(min int, vals ...int) int  {
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Println(max(1, 3, -1, 5))
	fmt.Println(min(2, 3, -1, 5))
}