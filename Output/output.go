package main
import ("fmt")

func main() {

  //Print() => prints it's arguments with their default format
  var i,j string = "Hello","World"

	fmt.Print(i,"\n")
	fmt.Print(j, "\n")
	fmt.Print(i, "\n", j)
	fmt.Print(i, " ", j)

  //Print() => A whitespace is added between the arguments, and a newline is added at the end
  fmt.Println(j,i)

  //Printf => formats its argument based on the given formatting verb and then prints them
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j) 
}