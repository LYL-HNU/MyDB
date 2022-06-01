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
	line, _, err := reader.ReadLine()
	if len(line) <= 0 {
		panic("Input wrong command " + err.Error())
	}
	input.Buffer = string(line)
	input.InputLength = len(input.Buffer)
	input.BufferLength = len(input.Buffer)
}

func (input *InputBuffer) CloseInput() {
	input.Buffer = ""
	input.InputLength = 0
	input.BufferLength = 0
}

func (input *InputBuffer) InitInput() {
	input.CloseInput()
}
