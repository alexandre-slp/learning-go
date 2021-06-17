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
		lastName: "Paes1",
		favIceCreamFlavors: []string{"strawberry", "vanilla"},
	}

	p2 := person{
		firstName: "Lele",
		lastName: "Paes2",
		favIceCreamFlavors: []string{"strawberry", "vanilla"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, mv := range m {
		fmt.Println(k, mv.firstName)
		fmt.Println("Favorite ice cream flavors:")
		for j, fv := range mv.favIceCreamFlavors {
			fmt.Println("\t", j, fv)
		}
	}
}
