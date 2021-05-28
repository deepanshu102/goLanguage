package main

import (
	"fmt"
)

type Dog struct {
	name   string
	age    int
	habits string
}

/*
 function with normal reciver
*/
func (variable Dog) normalFun() {
	fmt.Println("\nNormal Function Call")
}

/*
	Function with pointer reciver of type Struct
*/
func (a *Dog) init() {
	a.name = "Rocky"
	a.age = 0
	a.habits = ""
}
func (a *Dog) habit(hab string) {
	a.habits = hab
}

func main() {
	var dog Dog
	fmt.Printf("%+v", dog)
	dog.init()
	fmt.Printf("\n%+v", dog)
	dog.habit("Bark")
	fmt.Printf("\n%+v", dog)

	dog.normalFun()
}
