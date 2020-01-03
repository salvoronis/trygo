package msg

import(
  "testing"
)

type Messager interface {
  Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
  return m.Send("noc@example.com", "Critical Error", problem)
}

type MockMessage struct{
  email, subject string
  body []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error{
  m.email = email
  m.subject = subject
  m.body = body
  return nil
}

func TestingAlert(t *testing.T){
  msgr := new(MockMessage)
  body := []byte("Critical Error")
  Alert(msgr, body)

  if msgr.subject != "Critical Error" {
    t.Errorf("Expected ....")
  }
}
