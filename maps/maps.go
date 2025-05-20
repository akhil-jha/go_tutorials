package main

import (
	"fmt"
	"reflect"
)

func main() {
	myMap := make(map[string]string)
	myMap["India"] = "New Delhi"
	myMap["UK"] = "London"
	myMap["Czechia"] = "Prague"

	for key, value := range myMap {
		fmt.Printf("%v: %v\n", key, value)
	}

	for key := range myMap {
		fmt.Printf("%v\n", key)
	}

	for _, value := range myMap {
		fmt.Printf("%v\n", value)
	}
	fmt.Printf("\n%p\n", &myMap)

	fmt.Printf("Length of myMap: %v\n", len(myMap))

	delete(myMap, "UK")
	fmt.Printf("\n%p\n", &myMap)

	fmt.Println("Deleting  'UK'")
	fmt.Printf("Length of myMap: %v\n", len(myMap))

	clear(myMap)
	fmt.Printf("Length of myMap: %v\n", len(myMap))

	fmt.Printf("\n%p\n", &myMap)

	fmt.Println(reflect.TypeOf(myMap))

	myMultiplication := make(map[int]int)

	for i := 1; i < 11; i++ {
		myMultiplication[i] = i * 2
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("%v*2: %v\n", i, myMultiplication[i])
	} // Preserves the order

	fmt.Println("")

	for key, value := range myMultiplication {
		fmt.Printf("%v*2: %v\n", key, value)
	} // Does not preserve the order

}
