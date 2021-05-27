package main
import "fmt"
func main(){
age := 15
if age >= 16{
fmt.Print("Adult")
}else {
fmt.Print("Not an Adult")
}
if age >= 16{
fmt.Print("in School")
}else if  age >= 18 {
fmt.Println("in college")
}else{fmt.Println("probably dead")}

switch age{
case 16 : fmt.Println("Go Drive")
case 18 : fmt.Println("Go Vote")
default : fmt.Println("Go Have Fun")
}
}
