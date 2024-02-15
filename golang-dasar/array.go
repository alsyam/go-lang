package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "muhammad"
	names[1] = "al"
	names[2] = "syam"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	// panjang data
	fmt.Println(len(names))
	fmt.Println(len(values))
}
