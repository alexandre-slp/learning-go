package service

import "fmt"

type Speaker interface {
	Speak(string) string
}

type Sp struct{}

func (Sp) Speak(msg string) string {
	return fmt.Sprintf("hello %v", msg)
}
