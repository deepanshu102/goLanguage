package main

import (
	"fmt"
)

func main() {
	countries := map[string]string{
		"us": "USA",
		"fr": "France",
		"ch": "China"}
	fmt.Printf("%+v", countries)
	for k, v := range countries {
		fmt.Printf("\n%s->%s", k, v)
	}
	fmt.Printf("\n")
	deleteDemo(countries)
	fmt.Printf("From main -> %+v\n", countries)

}

//Delete function Demo
func deleteDemo(dic map[string]string) {
	fmt.Printf("Entered in DeleteDemo %+v\n", dic)
	delete(dic, "ch")
	fmt.Printf("After Delete China %+v\n", dic)

}
