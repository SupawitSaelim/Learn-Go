/*
var Can be used inside and outside of functions Variable 
declaration and value assignment can be done separately

:=  Can only be used inside functions
Variable declaration and value assignment cannot
be done separately (must be done in the same line)
*/

package main
import ("fmt")

var a int
var b int = 2
var c = 3

// d := 4 remove this comment and try to execute again
// d := also this

func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}