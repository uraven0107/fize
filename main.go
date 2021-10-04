package main

import (
	"fmt"
	"os"

	fize "github.com/uraven0107/fize/app"
)

func main() {
	if err := fize.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
