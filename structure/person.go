// this is person structure in go

package main

import "fmt"

type person struct {
	name string
	age  int
	tel  int
}

func main() {
	s := person{name: "John", age: 23, tel: 0123023}

	fmt.Printf("\t\vWelcome : %s you are %d years. Tel : %d\v\n", s.name, s.age, s.tel)
}
