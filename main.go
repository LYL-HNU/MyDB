package main

import (
	"MyDB/util"
	"fmt"
	"os"
	"strings"
)

func printPromt() {
	fmt.Printf("db > ")
}

func main() {
	input := util.InputBuffer{
		Buffer:       "",
		InputLength:  0,
		BufferLength: 0,
	}
	for true {
		printPromt()
		input.ReadInput()
		if strings.Compare(input.Buffer, ".exit") == 0 {
			input.CloseInput()
			os.Exit(0)
		} else {
			fmt.Println("Unknow command -> " + input.Buffer)
		}
	}
}
