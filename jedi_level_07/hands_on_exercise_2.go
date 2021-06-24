package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "James Bond",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
}

func changeMe(p *person) {
	p.first = "Alexandre"
}