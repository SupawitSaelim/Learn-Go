/*
Value Assignment After Declaration
It is possible to assign a value to a variable after it is declared. 
This is helpful for cases the value is not initially known.
*/

package main
import ("fmt")

func main() {
	var student1 string
	student1 = "Supawit"
	fmt.Println(student1)
}