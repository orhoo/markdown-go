package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var input = flag.String("i", "", "file to be parsed")

func main() {
	flag.Parse()

	fmt.Println(*input)

	if len(*input) == 0 {
		println("no markdown file specified")
		return
	}
	b, err := ioutil.ReadFile(*input)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(string(b))
}
