/*
The Printf() function first formats its argument based
on the given formatting verb and then prints them.
*/

// '%v' is used to print the value of the arguments
// '%T' is used to print the type of the arguments

package main
import ("fmt")

func main () {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}