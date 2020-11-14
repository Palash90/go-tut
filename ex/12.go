// Basic structs

package main

import "fmt"

type Product struct {
	// Any variable name with Init Cap can be accessed publicly
	Name string

	Quantity int

	Price float64

	Sell bool

	// Any variable name starting with small letter is treated as default modifier of Java
	shelf int
}

func main() {

	// Empty definition of struct, all values will be assigned default value for the corresponding type
	p1 := Product{}

	// The following will just print the values of the attributes in the struct
	fmt.Println(p1)
	// Use +v to better visualize the struct in a JSON like format
	fmt.Printf("%+v\n", p1)

	// Like Java/C/C# access variables in the struct
	p1.Name = "Diya"
	p1.Quantity = 100
	p1.Price = 2
	p1.Sell = true
	p1.shelf = 4

	fmt.Printf("%+v\n", p1)

	// Declaring and initializing the struct
	// Like assigning in map, you need the traling comma in each attribute
	p2 := Product{
		Name:     "Rangoli",
		Quantity: 50,
		Price:    10,
		Sell:     true,
		shelf:    5,
	}

	fmt.Printf("%+v\n", p2)

	// Declaring struct with only some attributes, attributes that are unassigned gets a default value
	p3 := Product{
		Name:     "Oil",
		Quantity: 50,
		Sell:     true,
	}

	fmt.Printf("%+v\n", p3)

}
