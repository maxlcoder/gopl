package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	summer := months[6:9]
	fmt.Println(summer)
	fmt.Println(summer[0])
	fmt.Println(summer[1:2])
	fmt.Println(summer[2:2])
	// fmt.Println(summer[:20])

	var runnes []rune

	for _, r := range "hello, 世界" {
		runnes = append(runnes, r)
	}
	fmt.Printf("%q\n", runnes)

	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Printf("%q", s[12])
	for i := 0; i < len(s); {
		fmt.Printf("%q", s[i:])
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Println(size)
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	var ages map[string]int

	fmt.Println(ages)
	fmt.Println(ages == nil)

	age, ok := ages["bob"]
	if !ok {
		fmt.Println("not ok")
		fmt.Println(age)
	}

	if age, ok := ages["bod"]; !ok {
		fmt.Println("not ok")
		fmt.Println(age)
	}
}
