package main

import (
	"errors"
	"fmt"
)

var (

	ValidationError = errors.New("Validation Error")
	NotFoundError = errors.New("Not Found Error")
)

func GetById(id string) error{
	if id == ""{
		return ValidationError
	} 

	if id != "Eko" {
		return NotFoundError
	}

	// sukses

	return nil
}

func main(){
	err := GetById("Eko")
	if err != nil {
		if errors.Is(err, ValidationError){
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError){
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}


}