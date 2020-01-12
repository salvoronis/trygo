package main

import(
  "net/http"
  "html/template"
)

func main(){
  http.HandleFunc("/", fileForm)
  http.ListenAndServe(":8080", nil)
}

func fileForm(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil)
  } else {
    mr, err := r.MultipartReader()
    if err != nil {
      panic("failed to read multipart message")
    }
    values := make(map[string][]string)
    maxValueBytes := int64(10 << 20)

    for {
      part, err := mr.NextPart()
      if err == io.EOF {
        break
      }

      name := part.FormName()
      if name == ""{
        continue
      }

      filename := part.FileName()
      var b bytes.Buffer
      if filename == ""{
        n, err := io.CopyN(&b, part, maxValueBytes)
        if err != nil && err != io.EOF {
          fmt.Fprint(w, "Error processing")
          return
        }
        maxValueBytes -= n
        if maxValueBytes == 0 {
          msg := "multipart message too large"
          fmt.Fprint(w, msg)
          return
        }
        values[name] = append(values[name], b.String())
        continue
      }
      dst, err := os.Create("./tmp/"+filename)
      defer dst.Close()
      if err != nil {
        return
      }
      for {
        buffer := make([]byte, 100000)
        cBytes, err := part.Read(buffer)
        if err == io.EOF {
          break
        }
        dst.Write(buffer[0:cBytes])
      }
    }
    fmt.Fprint(w, "Upload complite")
  }
}
