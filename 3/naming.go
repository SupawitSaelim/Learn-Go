/*
Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name 
(age, price, carname, etc.).

Go variable naming rules:

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords
*/

package main
import ("fmt")

//Camel Case
//Each word, except the first, starts with a capital letter:
var myVariableName string= "Camel"

//Pascal Case
var MyVariableName = "Pascal"

//Snake Case
var my_variable_name = "Snake"

func main() {
	fmt.Println(myVariableName,MyVariableName,my_variable_name)
}