package main

import "fmt"


func main() {
    div2 := func (i int) (int, bool) {
        return i /2, i % 2 == 0
    }
    
    for n:=1; n <= 10 ;n++ {
        fmt.Print(n, " ")
        fmt.Println(div2(n))
    }
}