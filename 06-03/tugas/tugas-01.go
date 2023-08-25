package tugas

import (
	"strconv"
	"strings"
)

type Anggota struct {
	Nama string
	Usia int
}

func Parser(input string) Anggota {
	dataAnggota := Anggota{}

	fields := strings.Split(input, ",")
	for _, field := range fields {
		parts := strings.Split(strings.TrimSpace(field), ":")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "nama":
			dataAnggota.Nama = value
		case "usia":
			usia, err := strconv.Atoi(value)
			if err == nil {
				dataAnggota.Usia = usia
			}
		}
	}

	return dataAnggota
}
