package main

import "fmt"

func main() {
	a := "This is my"
	b := "How are You"
	c := []byte(a)
	fmt.Println(string(a[1])) // taking out particular value
	fmt.Println(a + b)        //concatenation using plus operator
	fmt.Println(c)            // passing strings as byte format utf-8
}
