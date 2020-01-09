package main

import(
  "bytes"
  "net/smtp"
  "strconv"
  "text/template"
)

type EmailMessage struct{
  From, Subject, Body string
  To []string
}

type EmailCredentials struct{
  Username, Password, Server string
  Port int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}`

var t *template.Template

func init(){
  t = template.New("email")
  t.Parse(emailTemplate)
}

func main(){
  message := &EmailMessage{
    From: "salvoronis@gmail.com",
    To: []string{"mamekak@mail.ru"},
    Subject: "GAY NIGGERS FROM SPACE",
    Body: "welcome to the school boddy",
  }
  var body bytes.Buffer
  t.Execute(&body, message)

  authCreds := &EmailCredentials{
    Username: "salvoronis@gmail.com",
    Password: "************",
    Server: "smtp.gmail.com",
    Port: 465,
  }
  auth := smtp.PlainAuth("",authCreds.Username,authCreds.Password, authCreds.Server)
  smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),auth, message.From, message.To, body.Bytes())
}
