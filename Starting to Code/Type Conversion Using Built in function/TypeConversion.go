package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 12
	var foo = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", i, foo)

}
