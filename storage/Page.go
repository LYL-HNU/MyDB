package storage

const (
	PageSize = 4096
)

type Page struct {
	page [PageSize]byte
}

func (P *Page) GetRowsPerPage(R *Row) int {
	return PageSize / R.GetRowSize()
}

func (P *Page) GetDestinationRow(byteOffset uint32, row *Row) []byte {
	return P.page[byteOffset : byteOffset+uint32(row.GetRowSize())]
}
