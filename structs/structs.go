package main

import (
	"fmt"
)

type Employee struct {
	name     string
	age      int
	isRemote bool
	address  Address
}

type Address struct {
	city    string
	pincode int
}

func (e *Employee) updateName(newName string) { //e *Employee points to the actual class but e Employee only points to the local e variable.
	e.name = newName
	fmt.Println(e) // e has a local scope here when * is not used.

}

func (a *Address) updatePincode(newPincode int) {
	a.pincode = newPincode
}

func main() {
	address := Address{
		pincode: 201301,
		city:    "Noida",
	}
	emp1 := Employee{
		name:     "Jack",
		age:      35,
		isRemote: true,
		address:  address,
	}
	fmt.Println(emp1.name)

	job := struct {
		title  string
		salary int
	}{
		title:  "SW",
		salary: 100,
	}

	fmt.Println(job.salary)

	empPtr := &emp1 // pass by refernce
	empPtr.age = 200
	fmt.Println(emp1, *empPtr) // same output

	emp1.updateName("bob")
	fmt.Println(emp1.address.pincode)

	emp1.address.updatePincode(111122)

	fmt.Println(emp1)

}
