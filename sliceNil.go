package main
import "fmt"
func main(){
	var number []int
	printSlice(number)
	if number==nil{
		fmt.Printf("Slice is nil")
	}
}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
