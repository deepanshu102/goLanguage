package main
import (
	"fmt"
	"errors"
	"math"
)
func Sqrt(value float64)(float64,error){
	if(value<0){
		return 0,errors.New("Math:neagtive Number passed to Sqrt")
	}
	return math.Sqrt(value),nil
}
func main(){
	result,err:=Sqrt(-43)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
	result,err=Sqrt(9)
	if err!=nil{
		fmt.Println(err)
	}	else{
		fmt.Println(result)
	}
}