package main

import (
  "fmt"
)

func main(){
  defer goodbye() //отложенный вызов
  fmt.Println("hello")
}

func goodbye(){
  fmt.Println("goodbye")
}
