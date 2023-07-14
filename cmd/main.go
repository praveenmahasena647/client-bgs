package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/bg-client/api"
)

func main() {
	var err error = api.StartClient()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}
