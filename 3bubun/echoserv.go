package main

import(
  "bufio"
  "fmt"
  "net"
  "errors"
  "net/http"
)

func main(){
  //listen()
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
func listen(){
  listener, err := net.Listen("tcp", ":1488")
  if err != nil {
    fmt.Println("Failed to open port on 1488")
    return
  }
  for {
    conn, err := listener.Accept()
    if err != nil{
      fmt.Println("Error accepting connection")
      continue
    }
    go handle(conn)
  }
}
func handle(conn net.Conn){
  defer func(){
    if err := recover(); err != nil {
      fmt.Printf("Fatal error: %s\n", err) //теперь сервер не падает
    }
    conn.Close()
  }()

  reader := bufio.NewReader(conn)

  data,err := reader.ReadBytes('\n')
  if err != nil {
    fmt.Println("Failed to read from socket.")
    conn.Close()
  }
  response(data, conn)
}

/*func response(data []byte, conn net.Conn){
  defer func(){
    conn.Close()
  }()
  conn.Write(data)
}*/
func response(data []byte, conn net.Conn){
  conn.Write(data)
  panic(errors.New("Failure in response"))
}

func handler(res http.ResponseWriter, req *http.Request){
  panic(errors.New("Oh shit, here we go again"))
}
