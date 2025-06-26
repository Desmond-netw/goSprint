// this is person structure in go

package main

import "fmt"

// struct of person
type person struct {
	name    string
	age     int
	tel     int
	address Address
}

type Address struct {
	street int
	city   string
}

func main() {
	address := Address{city: "helsinki", street: 00400}
	person1 := person{name: "John", age: 23, tel: 024023523, address: address}

	fmt.Printf("\t\vWelcome : %s you are %d years. Tel : %d Address%v\v\n", person1.name, person1.age, person1.tel, person1.address)
}
