package main
import "fmt"
type Books struct{
	title string
	author string
	subject string
	book_id string
}

func main(){
	var Book1 Books
	var Book2 Books
	// Book1 specifications
	Book1.title="Go Progaming"
	Book1.author="Mahesh Kumar"
	Book1.subject="Go Programming Tutorials"
	Book1.book_id="7485961"

	//Book2 Specifications
	Book2.title="Black Book"
	Book2.author="Zara Ali"
	Book2.subject="Coding is life"
	Book2.book_id="7845123"
	fmt.Println(Book1) // Single-Shot Printing
	fmt.Printf("Bok1 title:-%s\n",Book1.title)
	fmt.Printf("Book1 author:-%s\n",Book1.author)
	fmt.Printf("Book1 subject:-%s\n",Book1.subject)
	fmt.Printf("Book1 book_id:-%s\n",Book1.book_id)
	fmt.Printf("%+v\n",Book2)
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   /* Its not Integer but still its running and printing 
   Value with Data type
   */
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id) 
}