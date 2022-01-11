package main
import ("fmt")

func main(){
	var arr1 = [3]int{1,2,3} //Length is defined
	arr2 := [5]int{4,5,6,7,8}
	arr3 := [...]string{"Test1","Test2"} //Length is inferred
	
	fmt.Println("Using var :", arr1)
	fmt.Println("Using := :", arr2)
	fmt.Println("Inferred length:", arr3)

	//Get elemets of specific index
	fmt.Println("Get first element:", arr2[0])
	fmt.Println("Get 2nd element:", arr3[1])

	//Changing the value of 1st index in arr3 array
	arr3[1] = "New Value" 
	fmt.Println("arr3 changed value:", arr3)

	//Array initialization
	a1 := [5]int{} //not initialized
	a2 := [5]int{1,2} //Partially initialized
	a3 := [5]int{1,2,3,4,5} //Fully initialized

	fmt.Println("Array a1 not initialized :", a1)
	fmt.Println("Array a2 is partially initialized := :", a2)
	fmt.Println("Array a3 is fully initialized:", a3)

	//Initializing only specific elements in an array
	a4 := [5]int{0:10,2:40,4:80}

	fmt.Println("Initializing specific elemet in array a4:", a4)

	//Get length of an array
	fmt.Println("Length of array a4", len(a4))


}