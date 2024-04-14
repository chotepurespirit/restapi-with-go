package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"` //สามารถใช้ json name เพื่อเปลี่ยนชื่อใน json ได้
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Earn", Age: 20},
	{ID: 2, Name: "Chote", Age: 20},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			log.Println("GET")
//เราไม่สามารถที่จะใส่ w.Write(users) ตรงๆ ได้เพราะมันเป็น type slice ไม่สามารถแปลงเป็น byte ได้
//ดังนั้นเราจะแปลง struct ให้เป็น json ก่อน
//โดยเราจะต้อง import package encoding/json มาใช้งาน
//json.Marshal จะทำการแปลง struct ให้เป็น json และ return ออกมาเป็น byte
			b, err := json.Marshal(users)
//ทุก err ต้อง return 200 ถ้าไม่มี err จะ return 500
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(b)
			return
		}
		if req.Method == "POST" {
			log.Println("POST")
			body, err := ioutil.ReadAll(req.Body) //ใช้ package ioutil ในการอ่าน body จาก request
			// ต่อจากนั้นให้ทำการ handle error ก่อน
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError) //ถ้ามี error ให้ return 500
				w.Write([]byte(err.Error()))                  //แสดง error ออกมา
				return
			}
			// จากนั้นให้ทำการแปลง body ที่ได้มาจาก request ให้เป็น struct
			var u User                     //สร้างตัวแปร u ขึ้นมาเพื่อเก็บค่าที่ได้จากการแปลง
			err = json.Unmarshal(body, &u) //ตรงนี้จะต้องส่ง pointer ไปเพื่อให้ json.Unmarshal แก้ไขค่าใน struct ได้
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) //ถ้ามี error ให้ return 400
				w.Write([]byte(err.Error()))         //แสดง error ออกมา
				return
			}

			users = append(users, u) //เพิ่มข้อมูลใหม่เข้าไปใน slice

			fmt.Fprintf(w, "hello %s created users", "POST") //แสดงข้อความออกไป
			return
		}

	})
	log.Println("Starting server on :2565") //แสดงข้อความว่า server กำลังทำงานอยู่ที่ port 2565
	log.Fatal(http.ListenAndServe(":2565", nil)) //คำสั่ง ListenAndServe จะทำงานแบบค้างเอาไว้ จนกว่าจะมีการปิด server
	log.Println("bye bye") //แสดงข้อความว่า server ถูกปิด
}
