/*
Go Variable Types
In Go, there are different types of variables, for example:

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false
*/

package main
import ("fmt")

func main() {
	var students1 string = "John" // type is string
	var students2 = "Jane" // type is inferred
	x := 2 // type is inferred

	fmt.Println(students1)
	fmt.Println(students2)
	fmt.Println(x)
}