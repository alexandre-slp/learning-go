package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james": {"`Shaken", "not stirred`", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range m {
		fmt.Println("key: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
