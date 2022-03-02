package main

import (
	"fmt"
	"log"
	"os"
)

//Struct for representation total slice
// first lvl of JSON object Parsing
type Users struct {
	Users []User `json: "users"`
}

//second lvl of JSON object Parsing
//Internal user representation
// docking json - name
type User struct {
	Name   string `json: "name"`
	Type   string `json: "type"`
	Age    int    `json: "age"`
	Social Social `json: "social"`
}

//third lvl of JSON object Parsing
//Social block representation
type Social struct {
	Vkontakte string `json: "vkontakte"`
	Facebook  string `json: "facebook"`
}

// 1 Deserialization
func main() {
	//create file descriptor
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfulyy")
}
