package main
import ("fmt")


func main() { 


	for i := 0; i < 5; i++ {
		if i == 3{
			fmt.Println(i)
		}else if i == 4{
			fmt.Println(i)
		}
	}


	for i := 0; i < 10; i++ {
		if i == 2{
			continue
			//break
		}
		fmt.Println(i)
	}

	//Nested loops
	quality := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange","mango"}


	for i:=0; i < len(quality); i++{
		for j:=0; j < len(fruits); j++{
			fmt.Println(quality[i],fruits[j])
		}
	}

	//Range
	for _, val := range fruits {
		fmt.Println(val)
	}

	

}