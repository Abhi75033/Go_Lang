package main

import (
	"bufio"
	"fmt"
	"os"
)

// The technique is for taking input from user is simillar to c lang

func main() {

	var Name string

	fmt.Println("Hey What is your name", Name)

	// fmt.Scan(&Name)

	// fmt.Println("Hello Mr.",Name)

	// But the drawback of scanf is it will scan the string or any datatypes  upto white space
	// so if you want to scan the string with space then you have to use the following method
	// reader mehod

	// reader is a variable and buffio

	

	reader := bufio.NewReader(os.Stdin)

	Name, _  = reader.ReadString('\n')

	/*

	    buffio = buffio is a standard package of Go lang 

		os.Stdin: Represents the standard input (usually the keyboard input from the user).

		bufio.NewReader(os.Stdin):

		Creates a new buffered reader that wraps around os.Stdin.
		Allows you to read input from the user efficiently, providing methods like ReadString, ReadBytes, and ReadLine to simplify input processing.
		reader :=:

		Assigns the newly created buffered reader to the variable reader so you can use it to read input elsewhere in your code.

		ReadString is a function that is used for reading the string and one agrument is passed "/n" that indicates that when to stop 
		or read upto "/n"
	*/


	fmt.Println("Hello Mr.",Name)

}
