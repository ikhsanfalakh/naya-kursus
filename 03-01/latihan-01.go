package main

import (
	"os"
	"runtime"
)

func main() {
	versi := runtime.Version()
	println(versi)

	cpu := runtime.NumCPU()
	println(cpu)

	os.Exit(0)
	dir := os.ReadDir()
	println(dir)
}
