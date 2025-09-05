package main
import ("fmt")

func main() {
	var a, b, c, d int = 1,3,5,7
	fmt.Println(a,b,c,d)


	//Note: If you use the type keyword, 
	//it is only possible to declare one type of variable per line.
	var e, f = 6, "Hello"
	g, h := 8, "World!"
	fmt.Println(e,g,f,h)

	//Go Variable Declaration in a Block
	var (
		z int
		x int = 1
		y string = "hello"

	)
	fmt.Println(z,x,y)
}
