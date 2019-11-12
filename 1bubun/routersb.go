package main

import(
  "path"
  "net/http"
  "strings"
  "fmt"
)

func main()  {
  pr := newPathResolver()
  pr.Add("GET /hello", hello)
  pr.Add("* /goodbye/*", goodbye)
  http.ListenAndServe("localhost:8080", pr)
}
func newPathResolver() *pathResolver  {
  return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
  handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc)  {
  p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
  check := req.Method + " " + req.URL.Path
  for pattern, HandlerFunc := range p.handlers {
    if ok, err := path.Match(pattern, check); ok && err == nil {
      HandlerFunc(res, req)
    } else if err != nil {
      fmt.Fprint(res, err)
    }
  }
  http.NotFound(res,req)
}
func hello(res http.ResponseWriter, req *http.Request)  {
  query := req.URL.Query()
  name := query.Get("name")
  if name == "" {
    name = "fucking semen"
  }
  fmt.Fprint(res, "little hand ", name)
}
func goodbye(res http.ResponseWriter, req *http.Request){
  path := req.URL.Path
  parts := strings.Split(path, "/")
  name := parts[2]
  if name == ""{
    name = "gay"
  }
  fmt.Fprint(res, "looox ", name)
}
