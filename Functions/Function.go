package main

import "fmt"

// main function is the special type of function in go loang the compiler starts executing code from here only

func main() {
	// main function
	fmt.Println("This is the main function we didn't required to call this function")
	fmt.Println("------------------------------------------------------------")
	// calling the message function
	message()
	fmt.Println("------------------------------------------------------------")
	// calling parametrized function with no return type
	addtion(20, 30)
	fmt.Println("------------------------------------------------------------")
	// calling parametrized function with  return type (1st way)

	ans := Multiply(20,10)

	fmt.Println("Multiplication of two nukmbers are",ans)

	fmt.Println("------------------------------------------------------------")

	// calling parametrized function with  return type (1st way)

	devide := Devide(20,10)

	fmt.Println("Multiplication of two nukmbers are",devide)
}

// Normal Function

func message() {
	fmt.Println("Hello, this is a normal function")
}

// parametrized function with no return type

func addtion(a, b int) {
	fmt.Println("Sum of two numbers is ", a+b)
}

// parametrized function with return type (1st way)

func Multiply(a, b int) (result int) {

	result = a * b

	return result

}

// parametrized function with return type (2nd way)

func Devide(a, b int) (result int) {

	return a/b

}
