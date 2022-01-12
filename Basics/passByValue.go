package main
import "fmt"

/*
When passing the variable to the fuction as an argument it makes a copy of that variable
In this fucntion the argument value is the copy of original value passed to the parameter.
So, the original value does not change.
*/
func updateName(x string) string{ 
	x = "wedge"
	return x
}

func  modifyBoolean(n bool) bool{
	return !n
}

//Original value of the argument value is changed
//Maps
func updateVehicle(y map[int]string) map[int]string{
	y[4] = "Helicopter"
	return y
}

//Arrays
func modifyArray(nations [3]string) [3]string{
	nations[1] = "Kathmandu"
	return nations
}

func main(){
	
	//Basic data type
	name := "tifa"

	name = updateName(name)

	fmt.Println(name)

	//#bool
	old := false
	fmt.Println("Before function call:", old)
	fmt.Println("After function call:", modifyBoolean(old))
	fmt.Println("After function call, value of old:", old)


	//Slices,maps,functions
	vehicles := map[int]string{
		1: "Car",
		2: "Bus",
		3: "Motar Bike",
	}

	// updateVehicle(vehicles)

	// fmt.Println("Updated Vehicles:", vehicles)
	fmt.Println("Before fucntion call:", vehicles)

	fmt.Println("Function call:", updateVehicle(vehicles)) //

	fmt.Println("After function call:", vehicles)

	//Arrays
	country := [3]string{"Nepal","Japan","Korea"}	
	fmt.Println("Before fucntion call:", country)

	fmt.Println("Function call:", modifyArray(country))

	fmt.Println("After function call:", country)
}
