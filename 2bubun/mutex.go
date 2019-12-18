package main

import(
  "bufio"
  "fmt"
  "os"
  "strings"
  "sync"
)

func main(){
  var wg sync.WaitGroup
  w := newWords()
  for _, f := range os.Args[1:]{
    wg.Add(1)
    go func(file string){
      if err := tallyWords(file,w); err != nil {
        fmt.Println(err.Error())
      }
      wg.Done()
    }(f)
  }
  wg.Wait()
  fmt.Println("Words that appear more than once:")
  w.Lock()
  defer w.Unlock()
  for word, count := range w.found {
    if count > 1 {
      fmt.Printf("%s: %d\n", word, count)
    }
  }
  //w.Unlock()
}

type words struct {
  sync.Mutex
  found map[string]int //извлекаемые слова помещаем в структуру
}

func newWords() *words {
  return &words{found: map[string]int{}} //создание нового экземпляра слова
}

func (w*words) add(word string, n int){ //фиксирует количество вхождений этого слова
  w.Lock()
  defer w.Unlock()
  count, ok := w.found[word]            //Если слово еще не зафиксировано, добавим его, в противном случае увеличим счетчик
  if !ok {
    w.found[word] = n
    return
  }
  w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error { //Открытие файла, анализ содержимого и подсчет найденных в нем слов. Функция копирования делает все необходимое
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan(){
    word := strings.ToLower(scanner.Text())
    dict.add(word, 1)
  }
  return scanner.Err()
}
