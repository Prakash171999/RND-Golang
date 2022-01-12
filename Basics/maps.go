package main
import (
	"fmt"
)

func main(){
	currency := map[string]float64{
		"Dollar": 125.50,
		"Pound": 155.50,
		"Yen": 175.50,
		"Riyal": 180.50,
	}	

	fmt.Println(currency)
	fmt.Println(currency["Dollar"])

	//looping maps
	for c, v := range currency {
		fmt.Println(c, "-", v)
	}

	//ints as key type
	phoneNumbers := map[int]string{
		9805898656: "Prakash",
		9810100125: "Sawal",
		9877854525: "Reyzan",
	}

	fmt.Println(phoneNumbers)
	fmt.Println(phoneNumbers[9805898656])
}