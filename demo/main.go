package main

import (
	"fmt"
	"os"

	yaclip "github.com/Sniper10754/yaclip/parser"
)

func main() {
	fmt.Println(os.Args)

	yaclip.Debug = true

	parser := yaclip.NewParserFromEnv()

	flags, err := parser.Parse()
	fmt.Println(flags)

	if err != nil {
		panic(err)
	}

}