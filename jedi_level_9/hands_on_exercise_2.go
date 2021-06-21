package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) speak(msg string) {
	fmt.Println(msg)
}

type human interface {
	speak(msg string)
}

func saySomething(h human)  {
	h.speak("something")
}

func main() {
	p := person{
		firstName: "Ale",
		lastName: "Paes",
	}

	//saySomething(p) wrong
	saySomething(&p)
	p.speak("hello")
}
