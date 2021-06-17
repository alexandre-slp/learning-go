package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck {
		vehicle: vehicle {
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}
	fmt.Println(t)
	fmt.Println(t.doors, t.color, t.fourWheel)

	s := sedan {
		vehicle: vehicle {
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println(s)
	fmt.Println(s.doors, s.color, s.luxury)
}
