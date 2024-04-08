package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID int `json:"id"` //สามารถใช้ json name เพื่อเปลี่ยนชื่อใน json ได้
	Name string `json:"name"`
	Age int `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Earn", Age: 20},
	{ID: 2, Name: "Chote", Age: 20},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
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
			body, err := ioutil.ReadAll(req.Body) //ReadAll จะทำการอ่าน body ทั้งหมด
			// ต่อจากนั้นให้ทำการ handle error ก่อน
			if err != nil {
	
			}
		}

		})
		/*check method
		if req.Method == "POST" {
			w.Write([]byte(`{"name" : "Chozzze" , "method" : "POST"}`))
			return
		}

		//writeHEADER
		w.WriteHeader(http.StatusMethodNotAllowed)
	})*/
	log.Println("Starting server on :2565")
	log.Fatal(http.ListenAndServe(":2565", nil)) //คำสั่ง ListenAndServe จะทำงานแบบค้างเอาไว้ จนกว่าจะมีการปิด server
	log.Println("bye bye")
}