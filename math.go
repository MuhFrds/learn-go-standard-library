package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40)) // dibulatkan ke atas
	fmt.Println(math.Floor(1.60)) // dibulatkan ke bawah
	fmt.Println(math.Round(1.60)) // dibulatkan ke yang terdekat
	fmt.Println(math.Max(10, 11)) // mencari yang paling besar
	fmt.Println(math.Min(10, 11)) // mencari yang paling kecil
}