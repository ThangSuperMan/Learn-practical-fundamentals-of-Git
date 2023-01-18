package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("hello")
	fmt.Println("modified")
	result := sum(1, 2)
	fmt.Println("result: ", result)
}
