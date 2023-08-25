package main

import (
	"0602/latihan"
	"fmt"
	"runtime"
	"sync"
)

func print(x int, str string, wg *sync.WaitGroup) int {
	defer wg.Done()
	i := 0
	for ; i < x; i++ {
		fmt.Println((i + 1), str)
	}

	return i
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	runtime.GOMAXPROCS(4)

	fmt.Println("Latihan-01")
	wg := sync.WaitGroup{}

	wg.Add(1)
	go print(100, "proses pertama", &wg)
	wg.Wait()

	wg.Add(1)
	go print(90, "proses kedua", &wg)

	fmt.Println("Latihan-02")
	// buatlah function yang terdapat variabel seperti di bawah ini
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	// lalu urutkan data phones tampilkan satu persatu per-detik dan sisipkan angka di depan masing-masing data sehingga menghasilkan output seperti ini (gunakan goroutine dan WaitGroup untuk mengerjakan soal ini):

	wg.Add(1)
	go latihan.PrintPhone(phones, &wg)
	wg.Wait()

	fmt.Println("selesai")
}
