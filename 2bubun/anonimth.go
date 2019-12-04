package main

import (
  "fmt"
  "runtime"
)

func main(){
  fmt.Println("outside a goruntime")
  go func() {
    fmt.Println("Inside a goruntime")
  }()
  fmt.Println("outside again")
  runtime.Gosched()
}
