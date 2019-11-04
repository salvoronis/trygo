package main

import (
  "log"
  "github.com/gotk3/gotk3/gtk"
)

func main(){
  gtk.Init(nil)

  win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
  if err != nil {
    log.Fatal("Не удалось создать окно", err)
  }
  win.SetTitle("rei")
  win.Connect("destroy", func(){
    gtk.MainQuit()
  })

  l, err := gtk.LabelNew("lol")
  if err != nil {
    log.Fatal("Не удалось создать окно", err)
  }

  //but, err := gtk.ButtonNewWithLabel("knopka")
  //if err != nil {
  //  log.Fatal("Не удалось создать окно", err)
  //}

  win.Add(l)
  //win.Add(but)

  win.SetDefaultSize(800,600)

  win.ShowAll()

  gtk.Main()
}
