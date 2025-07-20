package main

import (
	"fmt"
	"os"

	"github.com/driftsl/driftc/pkg/driftc"
)

func main() {
	var lexer driftc.Lexer
	tokens := must(lexer.Tokenize([]rune(string(must(os.ReadFile(os.Args[1])))), false))

	for i, token := range tokens {
		fmt.Printf("%d\t%+v\n", i, token)
	}

	var parser driftc.Parser
	result, err := parser.Parse(tokens)
	fmt.Printf("%+v\n", result)
	if err != nil {
		fmt.Println(err)
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
