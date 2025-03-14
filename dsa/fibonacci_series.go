// fibonacci series

package main

import "fmt"

func main() {
	var fib_series []int32 = []int32{0, 1}
	var n int = 9 - len(fib_series)
	// fmt.Printf("%v, %T \n", fib_series, fib_series)

	//for i, _ := range fib_series {
	for i := range n {

		sum := fib_series[i] + fib_series[i+1]
		fib_series = append(fib_series, sum)
		// fmt.Printf("%v, %v \n", i, v)

	}
	fmt.Printf("%v, %T \n", fib_series, fib_series)
	// fmt.Println(fib_series, n)

}
