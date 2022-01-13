package main

import ("fmt"
		"net/http"
		"io/ioutil"
)

const url = "https://lco.dev"

func main(){
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	//Majority of the response is read using ioutil
	databytes, err := ioutil.ReadAll(response.Body) 

	if err != nil {
		panic(err)
	}
	content := string(databytes) //Converting databyte to string
	fmt.Println(content)
}