package common

import (
	"MyDB/storage"
)

type Cursor struct {
	table   *storage.Table
	numRows uint32
	eof     bool
}

func (c *Cursor) GetNumRows() uint32 {
	return c.numRows
}

func (c *Cursor) TableStart(table *storage.Table) {
	c.table = table
	c.numRows = 0
	c.eof = table.GetNumRows() == 0
}

func (c *Cursor) TableEnd(table *storage.Table) {
	c.table = table
	c.numRows = table.GetNumRows()
	c.eof = true
}

func (c *Cursor) CursorAdvance() {
	c.numRows++
	if c.numRows == c.table.GetNumRows() {
		c.eof = true
	}
}

func (c *Cursor) IsEOF() bool {
	return c.eof
}

func (c *Cursor) GetTable() *storage.Table {
	return c.table
}
