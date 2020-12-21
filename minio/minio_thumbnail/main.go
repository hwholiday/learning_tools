package main

import (
	"os"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
