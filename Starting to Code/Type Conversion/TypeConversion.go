package main

import "fmt"

func main() {
	i := 12
	var foo = string(i)
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", i, foo)

}
