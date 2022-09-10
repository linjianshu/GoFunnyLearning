package main

import "fmt"

func main() {
	fmt.Printf("%-15v\n", "SpaceX")
	fmt.Printf("%15v\n", "Vigin Galactic")

	var (
		name = "ljs"
		age  = 15
	)
	var name1, age1 = "ljs", 10
	fmt.Println(name, age)
	fmt.Println(name1, age1)

}
