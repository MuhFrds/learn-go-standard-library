package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var StructField reflect.StructField = valueType.Field(i)
		fmt.Println(StructField.Name, "with type", StructField.Type)
		fmt.Println(StructField.Tag.Get("required"))
		fmt.Println(StructField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool){
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
	   f := t.Field(i)
	   if f.Tag.Get("required") == "true"{
		   data := reflect.ValueOf(value).Field(i).Interface()
		   result = data != ""
		   if result == false{
				return result
		   }
	   }
	}
	return result
}


func main() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "", ""})

	person := Person{
		Name: "Muh",
		Address: "Jakarta",
		Email: "test",
	}

	fmt.Println(isValid(person))
}