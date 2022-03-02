package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Struct for representation total slice
// first lvl of JSON object Parsing
type Users struct {
	Users []User `json:"users"`
}

//second lvl of JSON object Parsing
//Internal user representation
// docking json - name
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

//third lvl of JSON object Parsing
//Social block representation
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

//Func  for printf
func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social. VK: %s and FB: %s\n", u.Social.Vkontakte, u.Social.Facebook)
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

	//2 json to the Go
	//initializing the instance users
	var users Users
	//вычитываем содержимое jsonfile в виде последовательности байт
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	//From byteValue in users - deserialization
	json.Unmarshal(byteValue, &users)
	for _, u := range users.Users {
		fmt.Println("======================")
		PrintUser(&u)
	}

}
