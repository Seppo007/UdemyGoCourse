package main

import "fmt"

func main() {
	slice := newSliceOfInts()
	printEvenOrOdd(slice)
}

func newSliceOfInts() []int {
	var slice []int;
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}
	return slice
}

func printEvenOrOdd(slice []int) {
	for _, cur := range slice {
		fmt.Printf("%v is %v\n", cur, evenOdd(cur))
	}
}

func evenOdd(value int) string {
	if value%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
