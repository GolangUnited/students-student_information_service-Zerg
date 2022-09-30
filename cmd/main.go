package main

import "fmt"

// Comment

func main() {
	var (
		a = 40
		b = 1
	)
	fmt.Println("Result:", AddNumber(a, b))

}

func AddNumber(a int, b int) int {
	return a + b
}
