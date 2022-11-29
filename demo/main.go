package main

import (
	"fmt"
	"os"

	yaclip "github.com/Sniper10754/yaclip/parser"
)

func main() {
	fmt.Println(os.Args)

	yaclip.Debug = false

	parser := yaclip.NewParserFromEnv()

	flags, err := parser.Parse()
	
	if err != nil {
		panic(err)
	}

	for _, v := range flags {
		fmt.Printf("%s: %s -> %s\n", v.Type, v.Name, v.Value)
	}

}