package main

import (
	"0703/latihan"
	"0703/tugas"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	fmt.Println("Latihan 01")
	fmt.Println("GET http://localhost:8080/users")
	//http.HandleFunc("/users", latihan.Users)
	server.Get("/users", latihan.GetUsers)

	fmt.Println("GET http://localhost:8080/user?id=B001")
	//http.HandleFunc("/user", latihan.User)
	server.Get("/user", latihan.GetUser)

	url := "https://jsonplaceholder.typicode.com/users"

	fmt.Println("\nLatihan 02")
	fmt.Println("GET http://localhost:8080/users-new")
	//http.HandleFunc("/users-new", latihan.UsersNew(url))
	server.Get("users-baru", latihan.GetUsersBaru(url))
	server.Get("/users-new", func(c *fiber.Ctx) error {
		return latihan.GetUsersNew(url, c)
	})

	fmt.Println("GET http://localhost:8080/user-new?id=1")
	//http.HandleFunc("/user-new", latihan.UserNew(url))
	server.Get("/user-new", func(c *fiber.Ctx) error {
		return latihan.GetUserNew(url, c)
	})

	//soal 1
	fmt.Println("\ntugas 01")
	// buatlah API POST Nilai Mahasiswa untuk menambahkan data NilaiMahasiswa, ada beberapa ketentuan diantaranya:
	// 1. POST dapat menerima data dalam bentuk JSON maupun formData
	// 2. data yang di input hanya boleh Nama, MataKuliah dan Nilai saja, untuk ID dan IndeksNilai harus diolah terlebih dahulu baru di masukkan ke tambahkan ke NilaiMahasiswa
	// 3. Nilai hanya boleh diisi maksimal dengan angka 100
	// 4. untuk mengisi IndeksNilai memiliki kondisi: Nilai >= 80 indeksnya A, Nilai >= 70 dan Nilai < 80 indeksnya B, Nilai >= 60 dan Nilai < 70, indeksnya c Nilai >= 50 dan Nilai < 60 indeksnya D, Nilai < 50 indeksnya E
	fmt.Println("POST http://localhost:8080/nilaimahasiswa")
	server.Post("/nilaimahasiswa", tugas.PostNilaiMahasiswa)

	//soal 2
	fmt.Println("\ntugas 02")
	// buatlah API GET Nilai Mahasiswa untuk menampilkan semua data NilaiMahasiswa
	fmt.Println("GET http://localhost:8080/nilaimahasiswa")
	server.Get("/nilaimahasiswa", tugas.GetNilaiMahasiswa)

	fmt.Println("starting web server at http://localhost:8080/")
	//http.ListenAndServe(":8080", nil)
	server.Listen(":8090")
}
