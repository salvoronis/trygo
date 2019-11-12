package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

var help bool

func init(){
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
	flag.BoolVar(&help, "help", false, "hell to u")
}

func main(){
	flag.Parse()
	if help == true {
		flag.PrintDefaults()
	} else {
		if spanish == true {
			fmt.Printf("Hola %s!\n", *name)
		} else {
			fmt.Printf("Hello %s!\n", *name)
		}
	}
}