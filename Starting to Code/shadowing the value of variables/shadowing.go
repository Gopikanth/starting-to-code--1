package main

import "fmt"

var i = 32

func main() {
	//i := 12 //shadowing the value of i
	fmt.Printf("%v,%T", i, i)

}
