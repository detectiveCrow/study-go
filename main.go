package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("test!")
	fmt.Println("Hello world!")
	fmt.Println("----- Quote -----")
	fmt.Println(quote.Go())
	fmt.Println("----- ----- -----")
}