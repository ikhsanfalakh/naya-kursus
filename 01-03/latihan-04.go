package main

import "fmt"

/*
## Hari Ke-3: Pengambilan Keputusan
### Latihan-1
Buatlah sebuah program yang menentukan apakah seorang pahlawan dapat bertahan hidup atau tidak. Alkisah ada seorang pahlawan yang di tugaskan pergi ke sebuah kastil dimana kastil tersebut di kelilingi oleh banyak naga. Seekor naga dapat dibunuh dengan 2 peluru. Permasalahannya adalah pahlawan tersebut tidak tau berapa banyak naga yang ada dan berapa banyak peluru yang harus ia bawa. Jika pahlawan tersebut berhasil bertahan hidup kembalikan nilai true begitu pula sebaliknya
*/

func pahlawan(naga, peluru int) (statusPahlawan bool, sisaNaga int) {

	sisaNaga = peluru - naga*2
	if sisaNaga >= 0 {
		statusPahlawan = true
		sisaNaga = 0
	} else {
		statusPahlawan = false
	}

	return statusPahlawan, sisaNaga
}

func main() {
	naga := 10
	peluru := 21
	statusHidup, sisaNaga := pahlawan(naga, peluru)
	strHidup := ""

	if statusHidup {
		strHidup = "masih hidup"
	} else {
		strHidup = "telah gugur"
	}

	fmt.Println("Jumlah Naga ada", naga)
	fmt.Println("Pahlawan membawa peluru", peluru)
	fmt.Println("Naga yang berhasil dikalahkan", naga-sisaNaga)
	fmt.Println("Naga yang masih bertahan", sisaNaga)
	fmt.Println("Pahlawan", strHidup)
}
