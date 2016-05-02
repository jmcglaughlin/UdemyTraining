package main

import "fmt"

func wrapper() func() int {
    var x int
    return func() int {
        x++
        return x
    }
}

func main() { 
    incr1 := wrapper()
    incr2 := wrapper()
    
    fmt.Println(incr1())
    fmt.Println(incr1())
    fmt.Println(incr2())
    fmt.Println(incr2())
}