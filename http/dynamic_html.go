package main

import (
    "net/http"
    "text/template"
)
const doc = `
<!DOCTYPE html>
<html>
    <head>
      <title>  {{.Title}}</title>
    </head>
    <body>
        <h3>Hi, {{.Name}}. The fruits are:</h3>
        <ul>
            {{range .Fruit}}
                <li>{{.}}</li>
            {{end}}
        </ul>
    </body>
</html>
`
type Context struct {
    Title string
    Name  string
    Fruit [4]string
}
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Content Type", "text/html")
        // The template name "template" does not matter here
        templates := template.New("template")
        // "doc" is the constant that holds the HTML content
        templates.New("doc").Parse(doc)
        context := Context{
            Title: "My Fruits",
            Name: "John",
            Fruit: [4]string{"Apple", "Lemon", "Orange", "Banana"},
        }
        templates.Lookup("doc").Execute(w, context)
    })
    http.ListenAndServe(":8080", nil)
}
