package main

import "fmt"

func main() {
	i := 100
	a := &i
	*a = *a + 10
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n\n", *a)
	fmt.Printf("%v\n", &i)
	fmt.Printf("%v\n\n\n", i)

	c := &a
	**c = **c + 5100

	fmt.Println("After **C")
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n\n", &a)
	fmt.Printf("%v\n", *a)
	fmt.Printf("%v\n\n", &a)
	fmt.Printf("%v\n", **c)
	fmt.Printf("%v\n\n", &c)

	var myNum **int
	myNum = c
	fmt.Printf("%v\n\n", &myNum)
	fmt.Printf("%v\n\n", **myNum)

}
