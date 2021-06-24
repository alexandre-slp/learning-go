package main

import "fmt"

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	m := [][]string{s1, s2}

	for i, mVal := range m {
		fmt.Printf("line: %v \n", i)
		for j, inVal := range mVal {
			fmt.Printf("\t element: %v \t value: %v \n", j, inVal)
		}
	}
}