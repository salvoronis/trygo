package main

import(
  "fmt"
  "log"
  "log/syslog"
)

func main(){
  priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
  flags := log.Ldate | log.Lshortfile
  logger, err := syslog.NewLogger(priority, flags)
  if err != nil {
    fmt.Println("lol")
    return
  }
  logger.Println("syslog message")
}
