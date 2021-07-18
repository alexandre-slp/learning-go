package main

import (
	"fmt"
	"github.com/alexandre-slp/learning-go/jedi_level_13/quote"
	"github.com/alexandre-slp/learning-go/jedi_level_13/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}