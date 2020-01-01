package main

import(
  "log"
  "net"
)

func main(){
  conn, err := net.Dial("tcp", "localhost:1902")
  if err != nil {
    panic("oh shit")
  }
  defer conn.Close()
  logger := log.New(conn, "example ", log.Ldate|log.Lshortfile)
  logger.Println("regular")
  logger.Panicln("panic")
}
