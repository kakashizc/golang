package main

import "fmt"

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("苗苗")
}
func (d dog) speak() {
	fmt.Println("旺旺")
}

func main() {

}
