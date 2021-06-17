package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Ale",
		lastName: "Paes",
		favIceCreamFlavors: []string{"strawberry", "vanilla"},
	}

	p2 := person{
		firstName: "Lele",
		lastName: "Paes",
		favIceCreamFlavors: []string{"strawberry", "vanilla"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	fmt.Println("Favorite ice cream flavors:")
	for _, f := range p1.favIceCreamFlavors {
		fmt.Println("\t", f)
	}

	fmt.Println(p2.firstName, p2.lastName)
	fmt.Println("Favorite ice cream flavors:")
	for _, f := range p2.favIceCreamFlavors {
		fmt.Println("\t", f)
	}
}
