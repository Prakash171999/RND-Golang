package main
import ("fmt")

func main(){

//Single Variable Declaration
var a string
var b int
var c bool	
var test1 string = "John" //type is string
var test2 = "Jane" //type is inferred
x := 2 // type is inferred

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(test1)
fmt.Println(test2)
fmt.Println(x)

//Multiple Variable Declaration
var s, y, z, t int = 1, 3, 5, 7

  fmt.Println(s)
  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(t)

//Variable declaration in a block
var(
	q int
	w int = 1
	e string = "hello"
)

fmt.Println(q)
fmt.Println(w)
fmt.Println(e)

}