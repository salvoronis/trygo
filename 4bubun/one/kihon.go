package main

import(
  "html/template"
  "net/http"
)

type Page struct {
  Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request){
  p := &Page{
    Title: "An example",
    Content: "finally",
  }
  t := template.Must(template.ParseFiles("simple.html"))
  t.Execute(w, p)
}

func main(){
  http.HandleFunc("/", displayPage)
  http.ListenAndServe(":8080", nil)
}
