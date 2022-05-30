package util

import (
	"bufio"
	"os"
)

type InputBuffer struct {
	Buffer       string
	InputLength  int
	BufferLength int
}

func (input *InputBuffer) ReadInput() {
	reader := bufio.NewReader(os.Stdin)
	byte, _, err := reader.ReadLine()
	if len(byte) <= 0 {
		panic("Input wrong command " + err.Error())
	}
	input.Buffer = string(byte)
	input.InputLength = len(input.Buffer)
	input.BufferLength = len(input.Buffer)
}

func (input *InputBuffer) CloseInput() {
	input.Buffer = ""
	input.InputLength = 0
	input.BufferLength = 0
}
