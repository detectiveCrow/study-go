package main

import (
	"fmt"
)

var (
	testString1 string = "test 1"
	testString2 string = "test 2"
	testString3 string = "test 3"
	i           int    = 10 // Variable Shadowing
)

// var I int = 42 // for export

func main() {
	// option 1
	var i int
	i = 40
	// const i int = 10 // it's also work

	// option 2
	var j int = 50

	// option 3
	k := 100

	fmt.Printf("variable i is [ value = %v, type = %T ] \n", i, i)
	fmt.Printf("variable j is [ value = %v, type = %T ] \n", j, j)
	fmt.Printf("variable k is [ value = %v, type = %T ] \n", k, k)
}
