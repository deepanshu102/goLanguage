package main
import (
	"fmt"
	"time"
)

func counter(from,to int ,ch chan bool){
for i:=from;i<=to;i++{
	time.Sleep(50*time.Millisecond)
	fmt.Println(i)
}
ch<-true
}
func main(){
	ch:=make(chan bool)
	go counter(0,5,ch)
	go counter(6,10,ch)
	<-ch
}