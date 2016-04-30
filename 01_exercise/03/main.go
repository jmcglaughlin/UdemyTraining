package main

import (
	"fmt"
	"log"
)

func main() {
	var name string
    
    fmt.Print("Username: ")
	if _, err := fmt.Scan(&name); err != nil {
		log.Print("  Scan for name failed, due to ", err)
		return
	}

	fmt.Println("Hello", name)
}
