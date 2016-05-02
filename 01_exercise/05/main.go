package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if even(i) {
			fmt.Println(i)
		}
	}
}

func even(i int) bool {
	return i%2 == 0 
}
