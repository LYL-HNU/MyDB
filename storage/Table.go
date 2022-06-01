package storage

const (
	TableMaxPages = 100
)

type Table struct {
	numRows uint32
	pages   [TableMaxPages]*Page
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
