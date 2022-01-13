package main

import (
	"fmt"
	"encoding/json"
)

//Encoding JSON

type course struct {
	Name string `json:"coursename"` //Providing alias
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"` //Using - the password field will not be reflected who ever is consuming the api
	Tags []string `json:"tags,omitempty"` //omitempty says if the value is null, don't throw that field
}



func main(){
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in","abc123",[]string{"web-dev", "js"}},
		{"Go Bootcamp", 500, "LearnCodeOnline.in","abc123",[]string{"back-end", "js"}},
		{"Flutter Bootcamp", 350, "LearnCodeOnline.in","abc123",nil},
	}

	//package this data as JSON data

	// finalJson, err := json.MarshalIndent(lcoCourses, "lco", "\t")
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}