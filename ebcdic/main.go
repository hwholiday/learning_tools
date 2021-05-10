package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("/home/jk/projects/go/src/learning_tools/dd/EBCDIC.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Print(EncodeToString(data))
	fmt.Println(string(data))
}
