package main

import (
	"0603/latihan"
	"0603/tugas"
	"fmt"
	"sync"
)

func main() {
	//tugas-01
	fmt.Println("Tugas-01")

	input1 := "nama: joki siang sore, usia: 19"
	result1 := tugas.Parser(input1)
	fmt.Printf("Anggota{\n  Nama: \"%s\",\n  Usia: %d,\n}\n", result1.Nama, result1.Usia)
	fmt.Print(result1)

	input2 := "nama: Agung Febri, usia: 25"
	result2 := tugas.Parser(input2)
	fmt.Printf("Anggota{\n  Nama: \"%s\",\n  Usia: %d,\n}\n", result2.Nama, result2.Usia)

	//latihan-01
	fmt.Println("Latihan-01")

	movies := []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go latihan.GetMovies(movies, moviesChannel, &wg)

	go func() {
		wg.Wait()
		close(moviesChannel)
	}()

	for movie := range moviesChannel {
		fmt.Println("Received:", movie)
	}
}
