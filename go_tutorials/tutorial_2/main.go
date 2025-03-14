package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 23767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2

	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "hello" + " " + "world"
	fmt.Println(myString)
	fmt.Println(len("a")) // this bit len and not char len

	// to find len of the string
	fmt.Println(utf8.RuneCountInString("y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	// other ways to initiliaze the variable

	var intNum3 rune
	fmt.Println(intNum3)

	var myString1 = "Hi"
	fmt.Println(myString1)

	myVar := "text"
	fmt.Println(myVar)

	var var1, var2 int = 1, 2
	fmt.Println(var1, var2)

	var3, var4 := 1, 2
	fmt.Println(var3, var4)

	const myConst string = "my constant"
	fmt.Println(myConst)

}
