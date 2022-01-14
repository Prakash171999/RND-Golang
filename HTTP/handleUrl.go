package main

import ("fmt"
"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gsfsdfdsf"

func main(){
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	//Extracting params from url
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Paramas is:", val)
	}

	//Creating new URL
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:"lco.dev",
		Path:"/tutcss",
		RawPath: "user=prakash",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Creating Url: ", anotherUrl)
}