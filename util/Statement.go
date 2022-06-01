package util

import (
	"MyDB/storage"
)

type Statement struct {
	statementType int
	//插入行
	RowInsert storage.Row
}

func (Statement *Statement) SetStatementType(Type int) {
	Statement.statementType = Type
}

func (Statement *Statement) GetStatementType() int {
	return Statement.statementType
}
