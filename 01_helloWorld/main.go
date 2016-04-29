package main

import "fmt"

var hello = "Hello, 世界"

func main() {

	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)
}
