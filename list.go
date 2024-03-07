package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Muh")
	data.PushBack("Muhammad")
	data.PushBack("Firdaus")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Muh

	next := head.Next() // Muhammad
	fmt.Println(next.Value)

	next = next.Next() // Firdaus	
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}