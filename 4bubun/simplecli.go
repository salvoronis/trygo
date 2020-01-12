package main

import(
  "net/http"
  "time"
  "fmt"
  "io/ioutil"
  "os"
)

func main(){
  cc := &http.Client{Timeout: time.Second}
  res, err := cc.Get("http://goinpracticebook.com")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  b, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", b)
}
