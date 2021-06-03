package main
import (
	"html/template"
	"log"
	"net/http"
	"time"
)
type PageVaribles struct{
	Date string
	Time  string
}
func main(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func homePage( w http.ResponseWriter,r *http.Request){

	now:=time.Now()
	homePageVars:=PageVaribles{
		Date:now.Format("02-01-2006"),
	Time: now.Format("15:05:05"),
	}
	t,err:=template.ParseFiles("homepage.html")
	if err!=nil{
		log.Print("Template parsing error:-",err)
	}
	err=t.Execute(w,homePageVars)
	if err!=nil{
		log.Print("template executing error:-",err)
	}
}