package main

import(
  "errors"
  "fmt"
  "math/rand"
)

var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

var random = rand.New(rand.NewSource(35))

func main(){
  response, err := SendRequest("Hello")

  for err == ErrTimeout{
    fmt.Println("Timeout. Retrying.")
    response, err = SendRequest("hello")
  }
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response)
  }
}

func SendRequest(req string) (string, error){
  switch random.Int() % 3 {
  case 0:
    return "Seccess", nil
  case 1:
    return "", ErrRejected
  default:
    return "", ErrTimeout
  }
}
