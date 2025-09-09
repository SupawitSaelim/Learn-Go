
/*
Constant Rules
Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters 
(for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function
There are two types of constants:
Typed constants
Untyped constants
*/

package main
import ("fmt")

//Typed Constants
const A int = 1
//Untyped Constants
const B = 2
func main() {
	fmt.Println(A)
	fmt.Println(B)
}

// Note: In this case, the type of the constant is inferred from the value