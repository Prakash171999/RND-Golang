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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

	//Some cases where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is: %T \n", k, v, v)
	}
}
