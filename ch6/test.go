package main

type P struct {
	X int
	Y int
}

func (p *P) ScaleBy(factor float64)  {
	p.X = 1
	p.Y = 1
	return p.X * 1 + p.Y * 2
}

func main()  {
	p := &P{1, 2}

}
