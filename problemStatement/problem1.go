package main
import (
"fmt"
"sort")
func main(){
        list := []string{"Deepanshu","Alok","Rajan","Ramesh","Akilesh"}
	fmt.Println("Before Sorting-",list)
	sortList(list);	
}
func sortList(list [] string){
	sort.Strings(list)
	fmt.Println("After Sorting-",list)
}