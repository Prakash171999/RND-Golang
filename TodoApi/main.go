package main

import (
	"fmt"

	"./Config"
	"./Models"
	"./Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	//Creating a connection to DB
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status:", err)
	}

	defer Config.DB.Close()

	//run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup Routes
	r := Routes.SetupRouter()

	r.Run()
}
