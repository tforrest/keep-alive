package main

import (
	"flag"
	"fmt"
)

var jsonConfigFile = flag.String("file", "", "Config keep-alive via json")

func main() {
	flag.Parse()
	fmt.Println("Hello World")
	fmt.Printf("%s\n", *jsonConfigFile)
}
