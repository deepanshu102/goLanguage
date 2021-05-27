package main
import "fmt"
func main(){
	var a int =20
	var ip,k *int
	ip=&a
	fmt.Printf("Address of a variable:%x\n",&a)
	fmt.Printf("Address stored in ip variable: %x\n",ip)
	fmt.Printf("Value of *ip variable:%d\n",*ip)
	//k is nil pointer
	fmt.Printf("The value of ptr is :%x\n",k)
	if k==nil{
		fmt.Println("k variable is nil")
	}
	if ip!=nil{
		fmt.Println("Ip varibale is not nil")
	}
}