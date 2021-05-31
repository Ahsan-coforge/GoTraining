package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "Syed"
    s[1] = "Ahsan"
    s[2] = "Shahid"
    fmt.Println("Name:", s)
    fmt.Println("First Name:", s[1])

    fmt.Println("len:", len(s))

    s = append(s, "working")
    s = append(s, "in", "Coforge")
    fmt.Println("Full statement:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("Copied Statement:", c)

    l := s[2:6]
    fmt.Println("Slice 1:", l)

    l = s[:6]
    fmt.Println("Slice 2:", l)

    l = s[1:]
    fmt.Println("Slice 3:", l)

    t := []string{"in", "Greater", "Noida"}
    fmt.Println("Complete Statement:",l, t)

}