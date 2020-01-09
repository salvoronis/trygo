package main

import(
  "net/http"
)

func main(){
  dir := http.Dir("./file")
  handler := http.StripPrefix("/static", http.FileServer(dir))
  http.Handle("/static", handler)

  http.ListenAndServe(":8080", nil)
}
