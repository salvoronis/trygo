package main

import(
  "errors"
  "strings"
  "os"
  "fmt"
)

func Concat(parts ...string) (string, error){
  if len(parts) == 0 {
    return "", errors.New("No strings supplied") //есть ошибка
  }
  return strings.Join(parts, " "), nil //нет ошибок - error = nil
}

func main(){
  args := os.Args[1:]
  if result, err := Concat(args...); err != nil { //развертывание массива args, те args[0], [1]... и проверка на ошибку
    fmt.Printf("Error: %s\n", err)
  }else{
    fmt.Printf("Concated string: %s\n", result)
  }
}
