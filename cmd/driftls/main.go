package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/driftsl/driftc/pkg/driftls"
)

func main() {
	server := driftls.NewServer(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

	if err := server.Serve(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
