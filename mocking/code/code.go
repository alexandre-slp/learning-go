package main

import (
	"fmt"
	"github.com/alexandre-slp/learning-go/mocking/service"
)

func doSomething(msg string, speakerService service.Speaker) string {
	return speakerService.Speak(msg)
}
func main() {
	fmt.Println(doSomething("ale", service.Sp{}))
}
