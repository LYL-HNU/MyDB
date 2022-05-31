package main

import (
	"MyDB/util"
	"fmt"
	"os"
	"strings"
)

const (
	MetaCommandSuccess = iota
	MetaCommandUnrecognizedCommand
)

const (
	PrepareSuccess = iota
	PrepareUnrecognizedStatement
)

const (
	StatementInsert = iota
	StatementSelect
)

func printPromt() {
	fmt.Printf("db > ")
}

func doMetaCommand(inputBuffer *util.InputBuffer) int {
	if strings.Compare(inputBuffer.Buffer, ".exit") == 0 {
		inputBuffer.CloseInput()
		os.Exit(0)
	}
	return MetaCommandUnrecognizedCommand
}

func prepareStatament(statement *util.Statement, inputBuffer *util.InputBuffer) int {
	if strings.Compare(inputBuffer.Buffer[0:6], "insert") == 0 {
		statement.SetStatementType(StatementInsert)
		return PrepareSuccess
	}

	if strings.Compare(inputBuffer.Buffer, "select") == 0 {
		statement.SetStatementType(StatementSelect)
		return PrepareSuccess
	}
	return PrepareUnrecognizedStatement
}

func executeStatement(statement *util.Statement) {
	switch statement.GetStatementType() {
	case StatementInsert:
		fmt.Println("This is where we should do an insert")
		break
	case StatementSelect:
		fmt.Println("This is where we should do a select")
		break
	}
}

func main() {
	input := util.InputBuffer{
		Buffer:       "",
		InputLength:  0,
		BufferLength: 0,
	}
	var state util.Statement
	for true {
		printPromt()
		input.ReadInput()
		if strings.Compare(input.Buffer[0:1], ".") == 0 {
			switch doMetaCommand(&input) {
			case MetaCommandSuccess:
				//输入下一个命令
				continue
			case MetaCommandUnrecognizedCommand:
				fmt.Printf("Unknow command -> %s\n", input.Buffer)
				continue
			}
		}

		switch prepareStatament(&state, &input) {
		case PrepareSuccess:
			break
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", input.Buffer)
			continue
		}

		executeStatement(&state)
		fmt.Println("Executed.")
	}
}
