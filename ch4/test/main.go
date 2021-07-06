package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["t"] = 1

	age, ok := ages["s"]
	fmt.Println(ok)

	fmt.Println(age)

}
