package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)

  //Float
  var x float32 = 123.78
  var y float32 = 3.4e+38
  z := 3.3+36 //Type is not defined so by default type will be set to float64
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
  fmt.Printf("Type: %T, value: %v\n", z, z)
}