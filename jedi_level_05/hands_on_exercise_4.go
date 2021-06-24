package main

import "fmt"

func main() {
	as := struct {
		dict map[string]int
		list []string
	}{
		dict: map[string]int{
			"n1": 1,
			"n2": 2,
			"n3": 3,
		},
		list: []string{"abc", "def", "ghi"},
	}

	fmt.Println(as)
	for k, v := range as.dict {
		fmt.Println(k, v)
	}
	for i, v := range as.list {
		fmt.Println(i, v)
	}
}
