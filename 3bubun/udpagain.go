package main

import(
  "log"
  "net"
  "time"
)

func main(){
  timeout := 30 * time.Second
  conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
  if err != nil {
    panic("oh shit")
  }
  defer conn.Close()
  logger := log.New(conn, "example ", log.Ldate|log.Lshortfile)
  logger.Println("regular")
  logger.Panicln("panic")
}
