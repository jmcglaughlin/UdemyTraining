package main

import (
	"fmt"
	"log"
)

func main() {
	big := 1
    little := 1

	get("enter big num: ", &big)
	get("enter little num: ", &little)

	fmt.Println("the remainder of", big, "/", little, "is", big % little)

}
func get(q string, i *int) {

	fmt.Print(q)

	if _, err := fmt.Scan(i); err != nil {
		log.Print("  Scan for query failed, ", err)

	}
}
