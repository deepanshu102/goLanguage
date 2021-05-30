package main
import (
	"fmt"
)


func evenSum(from,to int,ch chan int){
	result:=0
	for i:=from;i<=to;i++{
		if i%2==0{
			fmt.Printf("%d\n",i)
			result+=i
		}
	}
	ch<-result
}
func SquareSum(from,to int,ch chan int){
	result:=0
	for i:=from;i<=to;i++{
		if i%2==0{
			fmt.Printf("\t%d\n",i)
			result+=i*i
		}
	}
	ch<-result
}
func main(){
	evenCh:=make(chan int)
	sqCh:=make(chan int)
	go evenSum(0,100,evenCh)
	go SquareSum(0,100,sqCh)
	fmt.Println(<-evenCh+<-sqCh)
}