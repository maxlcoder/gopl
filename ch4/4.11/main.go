package main

import "flag"

var (
	create = flag.Bool("c", false, "")
	list = flag.Bool("l", false, "")
	read = flag.Bool("r", false, "")
	edit = flag.Bool("e", false, "")

	owner = flag.String("owner", "", "")
	rep = flag.String("rep", "", "")
	number = flag.String("number", "", "")
	token = flag.String("token", "", "")


	title = flag.String("title", "", "")
	body = flag.String("body", "", "")

)

func main() {
	flag.Parse()
	switch {
	case *list:
		p := issu
		
	}
}
