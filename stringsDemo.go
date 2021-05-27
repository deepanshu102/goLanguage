package main
import (
	"fmt"
	"strings"
)

func main(){
	var greeting = "Hello World"
	fmt.Println("normal String");
	fmt.Printf("%s",greeting)
	fmt.Printf("\n")
	fmt.Printf("hex Bytes:")
	for i:=0;i<len(greeting);i++{
		fmt.Printf("%x",greeting[i])
	}
	fmt.Printf("\n")
	fmt.Println("String Length():-",len(greeting))

}