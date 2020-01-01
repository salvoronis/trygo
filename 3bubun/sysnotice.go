package main

import(
  "fmt"
  "log/syslog"
)

func main(){
  priority := syslog.LOG_LOCAL3
  logger, err := syslog.New(priority, "syssys")
  if err != nil {
    fmt.Println("lol")
    return
  }
  defer logger.Close()

  logger.Debug("debug message")
  logger.Notice("senpai")
  logger.Warning("warning")
  logger.Alert("alert")
}
