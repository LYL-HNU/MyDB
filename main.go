package main

import (
	"MyDB/storage"
	"MyDB/util"
	"encoding/binary"
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
	PrepareSyntaxError
)

const (
	StatementInsert = iota
	StatementSelect
)

const (
	ExecuteSuccess = iota
	ExecuteTableFull
)

func printPromt() {
	fmt.Printf("db > ")
}

func doMetaCommand(inputBuffer *util.InputBuffer, table *storage.Table) int {
	if strings.Compare(inputBuffer.Buffer, ".exit") == 0 {
		inputBuffer.CloseInput()
		table.InitTable()
		os.Exit(0)
	}
	return MetaCommandUnrecognizedCommand
}

func prepareStatament(statement *util.Statement, inputBuffer *util.InputBuffer) int {
	if strings.Compare(inputBuffer.Buffer[0:6], "insert") == 0 {
		statement.SetStatementType(StatementInsert)

		//获取用户输入id userName email
		var id uint32
		var userName, email string
		argsAssined, err := fmt.Sscanf(inputBuffer.Buffer, "insert %d %s %s", &id, &userName, &email)
		var idByte, userNameByte, emailByte []byte
		idByte = make([]byte, 4)
		userNameByte = make([]byte, storage.ColumnUsernameSize)
		emailByte = make([]byte, storage.ColumnEmailSize)
		binary.LittleEndian.PutUint32(idByte, id)
		userNameByte = []byte(userName)
		if cap(userNameByte) < storage.ColumnUsernameSize {
			userNameByte = append(userNameByte, make([]byte, storage.ColumnUsernameSize-cap(userNameByte))...)
		}
		emailByte = []byte(email)
		if cap(emailByte) < storage.ColumnEmailSize {
			emailByte = append(emailByte, make([]byte, storage.ColumnEmailSize-cap(emailByte))...)
		}
		//将id userName email 插入 Row 中
		statement.RowInsert.DeSerializeId(idByte)
		statement.RowInsert.DeSerializeUserName(userNameByte)
		statement.RowInsert.DeSerializeEmail(emailByte)
		if err != nil {
			panic(err)
		}
		if argsAssined < 3 {
			panic(fmt.Errorf("wrong input command -> %s", inputBuffer.Buffer))
		}
		return PrepareSuccess
	}

	if strings.Compare(inputBuffer.Buffer, "select") == 0 {
		statement.SetStatementType(StatementSelect)
		return PrepareSuccess
	}
	return PrepareUnrecognizedStatement
}

func executeStatement(statement *util.Statement, table *storage.Table) int {
	switch statement.GetStatementType() {
	case StatementInsert:
		return executeInsert(table, statement)
	case StatementSelect:
		return executeSelect(table)
	}
	return ExecuteSuccess
}

func serializeRow(row *storage.Row, destination []byte) {
	row.SerializeId(destination[row.GetIdOffset() : row.GetIdOffset()+row.GetIdSize()])
	row.SerializeUserName(destination[row.GetUserNameOffset() : row.GetUserNameOffset()+row.GetUserNameSize()])
	row.SerializeEmail(destination[row.GetEmailOffset() : row.GetEmailOffset()+row.GetEmailSize()])
}

func deserializeRow(row *storage.Row, destination []byte) {
	row.DeSerializeId(destination[row.GetIdOffset() : row.GetIdOffset()+row.GetIdSize()])
	row.DeSerializeUserName(destination[row.GetUserNameOffset() : row.GetUserNameOffset()+row.GetUserNameSize()])
	row.DeSerializeEmail(destination[row.GetEmailOffset() : row.GetEmailOffset()+row.GetEmailSize()])
}

func GetRowsPerPage(row *storage.Row) uint32 {
	const PageSize = 4096
	return uint32(PageSize / row.GetRowSize())
}

func rowSlot(table *storage.Table, rowNum uint32) []byte {
	var pageNum uint32
	pageNum = rowNum / GetRowsPerPage(new(storage.Row))
	var page *storage.Page
	page = table.GetPage(pageNum)
	if page == nil {
		//该页未有内容
		page = &storage.Page{}
	}
	rowOffset := rowNum % GetRowsPerPage(new(storage.Row))
	var r storage.Row
	byteOffset := rowOffset * uint32(r.GetRowSize())
	return page.GetDestinationRow(byteOffset, &r)
}

func executeInsert(table *storage.Table, statement *util.Statement) int {
	//假设目前一个表最多1000行
	if table.GetNumRows() >= 1000 {
		return ExecuteTableFull
	}
	row := &statement.RowInsert
	serializeRow(row, rowSlot(table, table.GetNumRows()))
	table.InsertRow()
	return ExecuteSuccess
}

func executeSelect(table *storage.Table) int {
	var row storage.Row
	numRows := int(table.GetNumRows())
	for i := 0; i < numRows; i++ {
		deserializeRow(&row, rowSlot(table, uint32(i)))
		row.PrintRow()
	}
	return ExecuteSuccess
}

func main() {
	input := util.InputBuffer{
		Buffer:       "",
		InputLength:  0,
		BufferLength: 0,
	}
	var table storage.Table
	table.InitTable()
	for true {
		var state util.Statement
		printPromt()
		input.InitInput()
		input.ReadInput()
		flag := strings.Compare(input.Buffer[0:1], ".")
		if flag == 0 {
			switch doMetaCommand(&input, &table) {
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
		case PrepareSyntaxError:
			fmt.Println("Syntax error. Could not parse statement.")
			continue
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", input.Buffer)
			continue
		}

		switch executeStatement(&state, &table) {
		case ExecuteSuccess:
			fmt.Println("Executed.")
			break
		case ExecuteTableFull:
			panic(fmt.Errorf("wrong. Table Full"))
		}
		fmt.Println("Executed.")
	}
}
