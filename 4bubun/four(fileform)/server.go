package main

import(
  "net/http"
  "html/template"
  "os"
  "fmt"
  "io"
)

func main(){
  http.HandleFunc("/", fileForm)
  http.ListenAndServe(":8080", nil)
}

func fileForm(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
  } else {
    f, h, err := r.FormFile("file")
    if err != nil {
      panic(err)
    }
    defer f.Close()
    filename := "./tmp/" + h.Filename
    out, err := os.Create(filename)
    if err != nil {
      panic(err)
    }
    defer out.Close()

    io.Copy(out,f)
    fmt.Fprint(w, "Upload Complite")
  }
}
