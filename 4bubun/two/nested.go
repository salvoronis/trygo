package main

import(
  "html/template"
  "net/http"
)

var t *template.Template

func init(){
  t = template.Must(template.ParseFiles("head.html", "main.html"))
}

type Page struct{
  Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request){
  p := &Page{
    Title: "An example",
    Content: "Have fun stormin' da castle",
  }
  t.ExecuteTemplate(w, "main.html", p)
}

func main(){
  http.HandleFunc("/", diaplayPage)
  http.ListenAndServe(":8080", nil)
}
