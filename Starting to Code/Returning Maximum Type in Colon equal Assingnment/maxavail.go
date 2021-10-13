package main

import "fmt"

func main() {
	var i float32 = 12.5
	j := 12.5 //  same value but takes maximum available size
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
}
