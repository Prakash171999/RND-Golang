package main
import ("fmt";"math")

//If inside a function
func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	age :=25

	if age >= 18{
		fmt.Println("You can ride car")
	}else if age < 18 && age > 16{
		fmt.Println("You can ride with a guardian")
	}else{
		fmt.Println("Sorry!You cannot ride")
	}

	num := 20

	if num >= 18{
		fmt.Println("You are adult")
		if age >= 18{
			fmt.Println("You can ride car")
		}
	}else{
		fmt.Println("You are child try again when you are more or equal to 18")
	}



fmt.Println(sqrt(5), sqrt(-4))

}