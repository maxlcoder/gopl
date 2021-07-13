package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	values := []int{1, 2, 3}
	fmt.Println(sum(values...))
}