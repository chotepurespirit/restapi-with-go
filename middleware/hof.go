package main

import "fmt"

type Decorator func(s string) error //type decorator เป็น function ที่รับ string และ return error

func Use(next Decorator) Decorator { //func Use จะ return Decorator ซึ่งเป็น function ที่รับ string และ return error
	return func(c string) error {
		fmt.Println("do something before")
		r := c + "should be green."
		return next(r)
	}
}

func home(s string) error {
	fmt.Println("home", s)
	return nil
}

func main() {
	wrapped := Use(home) //func home ถูกส่งเข้าไปใน func Use และ return ออกมาเป็น Decorator หรือพูดง่ายๆ คือ home จะแทนที่ next ใน func Use
	w := wrapped("world")
	fmt.Println("end result :", w)
}