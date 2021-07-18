package main

import (
	"fmt"
	"github.com/alexandre-slp/learning-go/mocking/service"
)

func doSomething(msg string) string {
	return service.Sp{}.Speak(msg)
}
func main() {
	fmt.Println(doSomething("ale"))
}
