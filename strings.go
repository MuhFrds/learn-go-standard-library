package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Muhammad Firdaus", "Muh"))
	fmt.Println(strings.Split("Muhammad Firdaus", " "))
	fmt.Println(strings.ToLower("Muhammad Firdaus"))
	fmt.Println(strings.ToUpper("Muhammad Firdaus"))
	fmt.Println(strings.Trim("     Muhammad Firdaus     ", " "))
	fmt.Println(strings.ReplaceAll("Muh Firdaus Muh Ahmad", "Muh", "Muhammad"))


}