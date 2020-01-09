package main

import(
  "net/http"
)

func main(){
  dir := http.Dir("./file")
  http.ListenAndServe(":8080", http.FileServer(dir))
}
