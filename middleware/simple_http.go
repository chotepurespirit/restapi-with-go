package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Chote", Age: 20},
}

func usersHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("POST")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		u := User{}
		err = json.Unmarshal(body, &u)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		users = append(users, u)
		fmt.Printf("% #v\n", users)

		fmt.Fprintf(w, "hello %s created users", "POST")
		return
	}

	if req.Method == "GET" {
		log.Println("GET")
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Server http middleware: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	})
}

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, req)
	log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		u, p, ok := req.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`can't parse the basic auth`))
			return 
		}
		if u != "apidesign" || p != "45678" {
			w.WriteHeader(http.StatusUnauthorized) //จะใช้เป็น 401 หรือ http.StatusUnauthorized ก็ได้
			w.Write([]byte(`username or password is incorrect`))
			return
		}
		fmt.Println("Auth passed.")
		next(w, req) //ถ้าผ่าน auth ให้ไปทำงานต่อที่ handler ถัดไป
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", AuthMiddleware(usersHandler))
	mux.HandleFunc("/health", AuthMiddleware(healthHandler))

	logMux := Logger{mux}

	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	log.Println("Server started at :2565")
	log.Fatal(srv.ListenAndServe())
	log.Println("bye bye!")
}