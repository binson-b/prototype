package main

import "fmt"

func main() {
	// Arrays
	// int32 store each int in 4 bytes
	// so array below is 12 bytes
	// default will be [0,0,0]
	var intArr [3]int32
	fmt.Println(intArr)
	intArr[1] = 123
	fmt.Println(intArr[1])

	// memory loc
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// other ways
	var intArr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr1)

	// Create an array with specific elements
	// composite array
	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice1 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice1...)
	fmt.Println(intSlice)

	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice2)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	delete(myMap2, "Adam")
	var age, ok = myMap2["Jasaon"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v\n", name, age)

	}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
