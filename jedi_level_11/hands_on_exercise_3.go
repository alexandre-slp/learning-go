package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func main() {
	e := customErr{info: "my test info"}
	foo(e)
}

func (ce customErr) Error() string {
	return fmt.Sprintf("my custom error: %v", ce.info)
}

func foo(e error)  {
	fmt.Println("FOO", e)
}
