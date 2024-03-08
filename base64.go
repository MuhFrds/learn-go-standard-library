package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Eko kurniawan khannedy"

	// konversi menjadi base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// mengembalikan menjadi string	
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

}