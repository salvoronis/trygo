package main

import(
  "html/template"
  "net/http"
)

var t = template.Must(template.ParseFiles("simple.html"))

type Page struct{
  Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request){
  p := &Page{
    Title: "An example",
    Content: "Have fun stormin' da castle",
  }
  t.Execute(w, p)
}

func main(){
  http.HandleFunc("/", diaplayPage)
  http.ListenAndServe(":8080", nil)
}
