package main

import "fmt"

func div2(i int) (int, bool) {
    return i /2, i % 2 == 0
}

func main() {
    n, even := div2(1)
    
    fmt.Println(n, even)
    
    for ;;n++ {
        fmt.Print(n, " ")
        fmt.Println(div2(n))
        if n == 10 {break}
    }
}