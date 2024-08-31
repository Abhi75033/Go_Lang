package Variable

import "fmt"

func Variables() {

	// First way of creating variables in GO lang var "Variable name " "data type" = "value"
	// Example

	var Name string = "Abhishek kumar"
	fmt.Println("Name:",Name)

	var age int = 23
	fmt.Println("age:",age)

	var height float64 = 12.23
	fmt.Println("Height:",height)

	var alcholic bool = true
	fmt.Println("Alcholic:",alcholic)

	// Second way to initialize variables in GO lang var "variable name" = "values"
	// Example
	fmt.Println("_______________________________________")
	var Department = "MCA"
	var Semester = 1
	var Matks = 100.00
	var passes = true

	fmt.Println("Department",Department)
	fmt.Println("Semester",Semester)
	fmt.Println("Matks",Matks)
	fmt.Println("passes",passes)

	fmt.Println("_______________________________________")


	// There is concept of const as well in the Go lang 
	// Example

	const pi = 3.14

	// We can't re-assign the value of const
	// pi = 3.15
	

	fmt.Println("pi:",pi)


	// Another way to Initialize variable in Go lang "variable name " := "value"
	// Example

	Collage := "Parul university "

	fmt.Println("Collage:",Collage)


	// Always remember if we want use any variable outside of this file then declare any variable with Uppercase
	// Example

	City := "Mugalsari"
	fmt.Println("City:",City)



}
