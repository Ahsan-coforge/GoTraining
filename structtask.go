package main

import "fmt"

type Details struct {
	Name    string
	city    string
	Pincode int
	Contact int
}

func main() {

	var a Details
	fmt.Println(a)

	a1 := Details{"Ahsan", "Delhi", 110006, 9953607212}

	fmt.Println("Details of first: ", a1)

	a2 := Details{Name: "Monika", city: "Dwarka",
		Pincode: 1122334, Contact: 9876543210}

	fmt.Println("Details of second: ", a2)

	a3 := Details{Name: "Alok",Pincode:1345634, Contact: 8976453120}
	fmt.Println("Details of third: ", a3)
}
