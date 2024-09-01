package main

import "fmt"

func main(){

	// 1st way Decleration or intalization of array in golang

	var Fruits[5] string
	Fruits[0]="Apple"
	Fruits[2] = "Mango"

	fmt.Println("The Array of Friuts is ",Fruits)

	// Accessing the elements of Array 

	fmt.Println("Accessing the elements at 0th Index and value is n",Fruits[0])
	fmt.Println("-----------------------------------------------------------------")

	// If we print an emty array then default value will print
	// Int default value is 0
	// Float default value is 0.0
	// String default value is ""
	// Boolean default value is false
	// Example

	var Marks[5] int
	fmt.Println("The Marks of Student is ",Marks)

	fmt.Println("-----------------------------------------------------------------")

	var Persons[5] string
	fmt.Println("The Persons is ",Persons)
	fmt.Printf("The Persons is %q\n",Persons)

	// %q is a fomart specifier, whenever we print any string so only value are getting printed without "" this qote
	// but using of %d helps us to print the valse as it is like "Abhishek"

	fmt.Println("-----------------------------------------------------------------")

	var Passed[5] bool
	fmt.Println("The Passed is ",Passed)

	fmt.Println("-----------------------------------------------------------------")

	var height[5] float64
	fmt.Println("The height is ",height)

	fmt.Println("-----------------------------------------------------------------")
}