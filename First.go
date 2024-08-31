// Because we are using the go outside of the ii's home directory then we have to import
// the ii package with the full path. using go mod init name ex: my learning

package main

import (
	"fmt"
	Variable "mylearning/Variables"
)

func main() {
	// fmt.Print("My name is abhishek kumar and i am learning go language\n")

	// myutil.Printmessage(10,20)

	Variable.Variables()

}

func addtion() {
	var a, b = 10, 20
	fmt.Println(a + b)
}
