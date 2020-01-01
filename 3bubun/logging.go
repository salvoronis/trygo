package main

import(
  "log"
  "os"
)

func main(){
  logfile,_ := os.Create("./log.txt")
  defer logfile.Close()

  logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

  logger.Println("regular message")
  logger.Fatalln("dye")
  logger.Println("regular message")
}
