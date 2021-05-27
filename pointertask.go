package main

import "fmt"

func main() {
    var name string = "ahsan"
    var pointer *string = &name

    fmt.Println("Name =", name)
    fmt.Println("Pointer =", pointer)

    fmt.Println("*pointer =", *pointer)

    *pointer = "Syed"
    fmt.Println("*pointer =", *pointer)
}