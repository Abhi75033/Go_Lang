package main

import "fmt"

func main() {
	fmt.Println(devide(20, 3))

	Devide, _ := devide(20,3)
	// "_" is used for ingoring any return value coming from from returning func 
	// Example Devide is returning two things (float64 and error) and we want to only float64 and ignore error then we use "_" this.

	fmt.Println("The devide of two number is ",Devide)

	//-----------------------------------------------------

	// if we want to handle the error so we can 

	Devide2,error:= devide2(10,0) 

	if error != nil{
		fmt.Println("Error is ",error)
	}else{
		fmt.Println("The devide of two number is ",Devide2)
	}
}

// ---------------------- 1st way to handle the error ------------------------- 

func devide(a, b float64) (float64, error) {
	/*
	In this function there are two types of returning values 
	1. float64 (Which is the result type of opearations between q,b)
	2. error (Which is the error type of operations between a,b. used for error handling)
	*/
	
	if b == 0 {
		return 0, fmt.Errorf("dominator can not be zero")
	}
	return a / b, nil
	// nill is used because we didn't have have any error this time so error = nill
}

// ---------------------- 2nd way to handle the error ------------------------- 

func devide2(a, b float64) (float64, error) {
	/*
	In this function there are two types of returning values 
	1. float64 (Which is the result type of opearations between q,b)
	2. error (Which is the error type of operations between a,b. used for error handling)
	*/
	
	if b == 0 {
		return 0, fmt.Errorf("dominator can not be zero")
	}
	return a / b, nil
	// nill is used because we didn't have have any error this time so error = nill
}
