package storage

const (
	PageSize = 4096
)

type Page struct {
	Page [PageSize]byte
}

func (P *Page) GetRowsPerPage(R *Row) int {
	return PageSize / R.GetRowSize()
}

func (P *Page) GetDestinationRow(byteOffset uint32, row *Row) []byte {
	return P.Page[byteOffset : byteOffset+uint32(row.GetRowSize())]
}

func (P *Page) SetPage(src []byte) {
	for i := 0; i < len(src); i++ {
		P.Page[i] = src[i]
	}
}
