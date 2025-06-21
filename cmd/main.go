package main

import (
	"fmt"
	"os"

	"github.com/driftsl/driftc/pkg/driftc"
)

func main() {
	lexer := driftc.Lexer{}
	tokens := must(lexer.Tokenize([]rune(string(must(os.ReadFile(os.Args[1]))))))

	for i, token := range tokens {
		fmt.Printf("%d\t%+v\n", i, token)
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
