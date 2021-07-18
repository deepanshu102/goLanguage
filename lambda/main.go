package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	s3session *s3.S3
)

const (
	REGION      = "ap-south-1"
	BUCKET_NAME = "lamdatestyoutube"
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Regin: aws.String(REGION)})))
}
func main() {
	lambda.Start(Handler)
}
func Handler(event InputEvent) (string, error) {
image:=GetImage(event.Link)
_,err:=s3session.PutObject(&s3.PutObjectInput{
	Body:bytes.NewReader(image),
	Bucker:aws.String(Bucket_NAME),
	key:aws.String(event.Key)
})
if err!=nil{
	return "something went wrong",err
}
return "Everthing worked YAY!",err
}

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func GetImage(url string) (bytes []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}
