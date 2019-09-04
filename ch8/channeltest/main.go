package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "A"
	fmt.Println(<-ch)
}
