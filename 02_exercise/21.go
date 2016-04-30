package main

import "fmt"


func main() {
    
    greatest := func (numbers ...int) (int) {
        var largest int
        for _, v := range numbers {
            if v > largest {
                largest = v
            }
        }
        return largest
    }
    
    max := greatest(1,2,315,4,5)

    fmt.Println(max)
    
    fmt.Println((true && false) || (false && true) || !(false && false))
}