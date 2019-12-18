package main

import(
  "compress/gzip"
  "fmt"
  "io"
  "os"
  "sync"
)

func main(){
  var wg sync.WaitGroup //нет необходимости инициализировать WaitGroup
  var i int = -1 //тк переменная i необходима за пределами цикла, она объявляется здесь
  var file string
  for i, file = range os.Args[1:]{
    wg.Add(1) //для каждого файла сообщить, что ожидается еще 1 операция

    go func (filename string){ //вызов функции и объявление о ее окончании
      compress(filename)
      wg.Done()
    }(file) //тк вызов сопрограммы происходит в цикле for требуется небольшая хитрость, чтобы передать параметр
  }
  wg.Wait() //прога ждет пока все сопрограммы не используют wg.Done
  fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
  in, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer in.Close()

  out, err := os.Create(filename+".gz")
  if err != nil {
    return err
  }
  defer out.Close()

  gzout := gzip.NewWriter(out)
  _, err = io.Copy(gzout, in)
  gzout.Close()

  return err
}
