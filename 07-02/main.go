package main

import (
	"0702/latihan"
	"0702/tugas"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Latihan 01")
	fmt.Println("GET http://localhost:8080/users")
	http.HandleFunc("/users", latihan.Users)
	fmt.Println("GET http://localhost:8080/user?id=B001")
	http.HandleFunc("/user", latihan.User)

	url := "https://jsonplaceholder.typicode.com/users"

	fmt.Println("\nLatihan 02")
	http.HandleFunc("/users-new", latihan.UsersNew(url))
	fmt.Println("GET http://localhost:8080/users-new")
	http.HandleFunc("/user-new", latihan.UserNew(url))
	fmt.Println("GET http://localhost:8080/user-new?id=1")

	fmt.Println("\nTugas 01")
	tugas.ShowData(url)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
