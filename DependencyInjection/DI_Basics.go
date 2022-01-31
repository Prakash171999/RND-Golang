package main

import "log"

//Implementation1
type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "User"
}

//Implementation2
type Greeter interface{
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string){
	log.Printf("Hello %s!! Nice to meet you", userName)
}

//Implementation3
type PopularDatabase struct{}

func (db PopularDatabase) GetUser() string {
	return "Will Smith"
}

func (db PopularDatabase) GetAllUsers() []string{
	return []string{}
}


type Program struct {
	Db		Database
	Greeter	Greeter
}

func (p Program) Execute(){
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main(){
	// db := DefaultDatabase{}
	db := PopularDatabase{}
	greeter := NiceGreeter{}
	p := Program{Db: db,Greeter: greeter}

	p.Execute()
}