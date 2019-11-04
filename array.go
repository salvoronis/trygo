package main

import (
  "fmt"
)

var sum int = 0

func main(){
  ar := [4][4]int{
    {-1,2,3,4},
    {-1,2,3,4},
    {-1,2,3,4},
    {-1,2,3,4},
  }
  newarr := FirstTask(ar)
  fmt.Println(ar)
  fmt.Println(newarr)
  //ТУТ НАЧИНАЕТСЯ 2 ЗАДАНИЕ
  vtoroe := DiagonalSum(ar)
  fmt.Println(vtoroe)
  //3
  a := [10]int{1,2,3,4,5,6,7,8,9,10}
  b := [10]int{1,2,3,4,5,6,7,8,9,10}
  tretie := SEMEN(a,b)
  fmt.Println(tretie)
}
//Эта функция для 1 задания
func FirstTask(array [4][4]int) ([4][4]int){
  for i := 0; i < 4; i++ {
    for j := 0; j < 4; j++ {
      if array[i][j] > 0 {
        sum++
      }
    }
  }
  for i:= 0; i < 4; i++{
    array[i][i] = array[i][i] + sum
  }
  return array
}
//Эта функция для 2 задания
func DiagonalSum(array [4][4]int) (int){
  var summa int = 0
  for i := 0; i < 4; i++ {
    summa = array[i][i] + summa
  }
  for i := 3; i >= 0; i-- {
    summa = array[3-i][i] + summa
  }
  return summa
}
//3
func SEMEN (a [10]int, b [10]int)([20]int){
  var vivod [20]int
  for i := 0; i < 20; i++{
    if i%2 == 0 {
      vivod[i]=a[i/2]
    } else {
      vivod[i]=b[i/2]
    }
  }
  return vivod
}
