package main

import(
  "errors"
)

func main(){
  panic(errors.New("Oops, I did it again")) //panic(interface{})
}
