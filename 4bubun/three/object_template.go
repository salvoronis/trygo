package main

import(
  "bytes"
  "html/template"
  "net/http"
)

var t *template.Template
var qc template.HTML

func init(){
  t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct{
  Title string
  Content template.HTML
}

type Quote struct{
  Quote, Person string
}

func main(){
  q := &Quote{
    Quote: `My cock is much bigger than yours. My shit smells better than yours.`,
    Person: "System Of A Down",
  }
  var b bytes.Buffer
  t.ExecuteTemplate(&b,"quote.html",q)
  qc = template.HTML(b.String())
  http.HandleFunc("/", diaplayPage)
  http.ListenAndServe(":8080", nil)
}

func diaplayPage(w http.ResponseWriter, q *http.Request){
  p := &Page{
    Title: "A User",
    Content: qc,
  }
  t.ExecuteTemplate(w, "index.html", p)
}
