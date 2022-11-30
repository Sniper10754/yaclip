package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var Debug bool = false

var ErrorParse error = errors.New("yaclip: invalid syntax")

type FlagType string

const VAR FlagType = "VAR"
const OPT FlagType = "OPT"

func PrintOnDebug(a ...any) {
	if Debug {
		fmt.Println(a...)
	}
}

type Parser struct {
	Args         []string
	PanicOnError bool
}

type Flag struct {
	Name  string
	Value string
	Type  FlagType
}

func NewParserFromEnv() Parser {
	return NewParser(os.Args)
}

func NewParser(args []string) Parser {
	return Parser{
		Args: args,
	}
}

func (p Parser) HandleError(err error) error {
	if p.PanicOnError {
		panic(err)
	} else {
		return err
	}
}

func (p Parser) Parse() ([]Flag, error) {
	flags := []Flag{}

	for i, arg := range p.Args {
		if i == 0 {
			continue
		}

		if !strings.HasPrefix(arg, "--") {
			PrintOnDebug("error: detected invalid syntax")

			return nil, p.HandleError(errors.New(ErrorParse.Error() + "; invalid argument " + arg))
		}

		if strings.Contains(arg, "=") {
			z := strings.Split(arg, "=")

			if z[1] == "" {
				return nil, p.HandleError(errors.New(ErrorParse.Error() + "; no argument found after '=' "))
			}

			flags = append(flags, Flag{
				Name:  z[0],
				Value: z[1],
				Type:  VAR,
			})
		} else {
			flags = append(flags, Flag{
				Name: arg,
				Type: OPT,
			})
		}

	}

	PrintOnDebug("Parsing process finished, now parser flags is", flags)
	return flags, nil
}
