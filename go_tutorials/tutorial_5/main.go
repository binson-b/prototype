package main

import "fmt"

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
}
