package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int
	Name string
	Age int
}

func main() {
	data := []byte(`{
	"id":1,
	"name":"Earnnie",
	"age":20
	}`)

	var u User
	
	err := json.Unmarshal(data, &u) //ตรงนี้จะต้องส่ง pointer ไปเพื่อให้ json.Unmarshal แก้ไขค่าใน struct ได้, 
									//ถ้าอยากให้ key รับค่าต้องใช้ตัวพิมพ์ใหญ่บน struct ด้วย
	fmt.Printf("% #v\n", u)
	fmt.Println(err)
}