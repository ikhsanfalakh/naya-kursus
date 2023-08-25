package main

import (
	"fmt"
)

/*
Buatlah program yang menentukan di tahun ke berapa sebuah kota akan mencapai atau melebihi kapasitas penduduknya. Dengan
ketentuan p0 adalah jumlah penduduk saat ini, percent adalah persentase kenaikan jumlah penduduk pertahun, aug adalah
jumlah pendatang baru pertahun, dan p adalah kapasitas dari kota tersebut.

At the end of the first year there will be:
1000 + 1000 * 0.02 + 50 => 1070 inhabitants

At the end of the 2nd year there will be:
1070 + 1070 * 0.02 + 50 => 1141 inhabitants (** number of inhabitants is an integer **)

At the end of the 3rd year there will be:
1141 + 1141 * 0.02 + 50 => 1213

It will need 3 entire years.
*/

func YearsToReachPopulation(p0, percent, aug, p int) int {
	population := p0
	years := 0

	for population < p {
		population = population + int(float64(population)*float64(percent)/100) + aug
		years++
	}

	return years
}

func main() {
	p0 := 200000  // Jumlah penduduk saat ini
	percent := 20 // Persentase kenaikan penduduk per tahun
	aug := 5000   // Jumlah pendatang baru per tahun
	p := 1000000  // Kapasitas penduduk kota

	years := YearsToReachPopulation(p0, percent, aug, p)
	fmt.Printf("It will need %d entire years.\n", years)
}
