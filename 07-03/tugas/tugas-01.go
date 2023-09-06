package tugas

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

var allNilaiMahasiswa = []NilaiMahasiswa{}

func GetNilaiMahasiswa(c *fiber.Ctx) error {
	// nilaimhs := allNilaiMahasiswa
	// dataNilaimhs, err := json.Marshal(nilaimhs)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(dataNilaimhs)
	return c.JSON(fiber.Map{"data": allNilaiMahasiswa})
}

func PostNilaiMahasiswa(c *fiber.Ctx) error {
	var nilaimhs NilaiMahasiswa

	if len(allNilaiMahasiswa) > 0 {
		nilaimhs.ID = allNilaiMahasiswa[len(allNilaiMahasiswa)-1].ID + 1
	} else {
		nilaimhs.ID = 1
	}

	//if c.Method() == "POST" {
	if c.Get("Content-Type") == "application/json" {
		// parse dari json
		if err := c.BodyParser(&nilaimhs); err != nil {
			log.Fatal(err)
		}
	} else {
		// parse dari form
		nama := c.FormValue("nama")
		matakuliah := c.FormValue("mata_kuliah")
		getNilai := c.FormValue("nilai")
		nilai, _ := strconv.Atoi(getNilai)
		id := nilaimhs.ID

		if psn, stt := cekNilai(nilai); stt == false {
			return c.Status(fiber.StatusBadRequest).SendString(psn)
		}

		nilaimhs = NilaiMahasiswa{
			ID:         id,
			Nama:       nama,
			MataKuliah: matakuliah,
			Nilai:      uint(nilai),
		}
	}

	if psn, stt := cekNilai(int(nilaimhs.Nilai)); stt == false {
		return c.Status(fiber.StatusBadRequest).SendString(psn)
	}

	nilaimhs.IndeksNilai = getIndeks(int(nilaimhs.Nilai))
	allNilaiMahasiswa = append(allNilaiMahasiswa, nilaimhs)
	//dataJsonNilaiMhs, _ := json.Marshal(nilaimhs) // to byte
	return c.JSON(fiber.Map{"data": allNilaiMahasiswa}) //c.Send(dataJsonNilaiMhs)               // kirim sebagai respons JSON
	//}

	//return c.Status(fiber.StatusNotFound).SendString("NOT FOUND")
}

func cekNilai(nilai int) (string, bool) {
	var pesan string
	var status bool
	if nilai > 100 {
		pesan = "Nilai Maximum must be 100"
		status = false
	} else if nilai < 0 {
		pesan = "Nilai Minimum must be 0"
		status = false
	} else {
		pesan = "Nilai Benar"
		status = true
	}
	return pesan, status
}

func getIndeks(nilai int) string {
	var indeksnya string
	if nilai >= 80 {
		indeksnya = "A"
	} else if nilai >= 70 && nilai < 80 {
		indeksnya = "B"
	} else if nilai >= 60 && nilai < 70 {
		indeksnya = "C"
	} else if nilai >= 50 && nilai < 60 {
		indeksnya = "D"
	} else if nilai < 50 {
		indeksnya = "E"
	}
	return indeksnya
}
