package main
import ("fmt"
"math")

//Note: Like struc is the collection of fields. interface is the collection of methods
//If a function has a print() method then one can implement it however one wants it. It will satisfy the interface as long as the function contains all the methods of the interface

type rect struct{
	width float32
	length float32
}

type circle struct{
	radius float32
}

type geometry interface{
	area() float32
	perim() float32
}

func (c circle) area() float32{
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 *math.Pi * c.radius
}

func (r rect) perim() float32 {
	return 2 * r.width * 2 * r.length
}

func (r rect) area() float32 {
	return r.width * r.length
}

func measure(g geometry){
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


func main(){
	myCircle := circle{radius:5}

	myRect := rect{width:20, length:20}

	// fmt.Println(myCircle.area())
	// fmt.Println(myRect.area())

	measure(myCircle)
	measure(myRect)
}