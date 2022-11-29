package main

import (
	"fmt"

	yaclip "github.com/Sniper10754/yaclip/parser"
)

func main() {
	yaclip.Debug = true

	parser := yaclip.NewParserFromEnv()

	flags, err := parser.Parse()
	
	if err != nil {
		panic(err)
	}

	for _, v := range flags {
		fmt.Printf("%s: %s -> %s\n", v.Type, v.Name, v.Value)
	}

}