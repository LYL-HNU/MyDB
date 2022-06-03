package storage

import (
	"fmt"
	"os"
)

const (
	TableMaxPages = 100
)

type Table struct {
	numRows    uint32
	pages      [TableMaxPages]*Page
	bufferpool *BufferPool
}

func (Table *Table) GetPage(pageNum uint32) *Page {
	return Table.pages[pageNum]
}

// InitTable 初始化或清楚
func (Table *Table) InitTable() {
	Table.numRows = 0
	for i := 0; i < TableMaxPages; i++ {
		Table.pages[i] = nil
	}
}

// GetNumRows 返回表中行数
func (Table *Table) GetNumRows() uint32 {
	return Table.numRows
}

// InsertRow 表中新增了一行
func (Table *Table) InsertRow() {
	Table.numRows++
}

func (Table *Table) SetPage(pageNum uint32, page *Page) {
	Table.pages[pageNum] = page
}

func (Table *Table) bufferPoolOpen(fileName string) {
	//不存在就创建文件，并在尾部写入
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Printf("open wrong database file. err: %s \n", err.Error())
		os.Exit(1)
	}
	Table.bufferpool.SetFileDesc(file)
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("open wrong database file")
		os.Exit(1)
	}
	Table.bufferpool.SetFileLength(uint32(fileStat.Size()))
}

func (Table *Table) DBOpen(fileName string) {
	Table.InitTable()
	Table.bufferpool = new(BufferPool)
	Table.bufferpool.InitBufferPool()
	Table.bufferPoolOpen(fileName)
	Table.numRows = Table.bufferpool.GetFileLength() / uint32(new(Row).GetRowSize())
	Table.bufferpool.InitBufferPool()
}

func (Table *Table) DBClose() {
	numFullPages := Table.GetNumRows() / uint32(new(Row).GetRowSize())
	for i := 0; i < int(numFullPages); i++ {
		if Table.bufferpool.pages[i] == nil {
			//空页
			continue
		}
		Table.bufferpool.FlushBack(uint32(i))
		Table.bufferpool.pages[i] = nil
	}
	rowsPerPage := PageSize / new(Row).GetRowSize()
	additionalRows := Table.numRows % uint32(rowsPerPage)
	//有部分行数不足一页
	if additionalRows > 0 {
		pageNum := numFullPages
		Table.bufferpool.FlushBack(pageNum)
		Table.bufferpool.pages[pageNum] = nil
	}
}
