package util

type Statement struct {
	statementType int
}

func (Statement *Statement) SetStatementType(Type int) {
	Statement.statementType = Type
}

func (Statement *Statement) GetStatementType() int {
	return Statement.statementType
}
