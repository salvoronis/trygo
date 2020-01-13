package main

import(
  "net/http"
  "encoding/json"
  "fmt"
)

type Error struct{
  HTTPCode int `json:"-"`
  Code int `json:"code,omitempty"`
  Message string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error){
  data := struct {
    Err Error `json:"error"`
  }{e}

  b,err := json.Marshal(data)
  if err != nil {
    http.Error(w, "Internal Server Error", 500)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(e.HTTPCode)
  fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request){
  e := Error{
    HTTPCode: http.StatusForbidden,
    Code: 123,
    Message: "Oh shit",
  }
  JSONError(w,e)
}

func main(){
  http.HandleFunc("/", displayError)
  http.ListenAndServe(":8080", nil)
}
