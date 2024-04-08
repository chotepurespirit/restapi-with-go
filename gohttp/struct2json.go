package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID:  1,
		Name: "Earnnie",
		Age: 20, //json.Marshal จะทำการแปลง struct ให้เป็น json และ return ออกมาเป็น byte
	}
	b, err := json.Marshal(u)
	fmt.Printf("byte : %T \n", b) //%T เป็นการแสดง type ของตัวแปร
	fmt.Printf("byte : %s \n", b)
	fmt.Println(err)
}
