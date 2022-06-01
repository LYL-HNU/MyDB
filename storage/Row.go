package storage

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const (
	// ColumnUsernameSize userName最大字节
	ColumnUsernameSize = 32
	// ColumnEmailSize email最大字节
	ColumnEmailSize = 255
)

type Row struct {
	id       int32
	userName [ColumnUsernameSize]byte
	email    [ColumnEmailSize]byte
}

func (R *Row) PrintRow() {
	fmt.Printf("(%d, %s, %s)\n", R.id, R.userName, R.email)
}

func (R *Row) GetIdSize() int {
	return int(unsafe.Sizeof(R.id))
}

func (R *Row) GetUserNameSize() int {
	return int(unsafe.Sizeof(R.userName))
}

func (R *Row) GetEmailSize() int {
	return int(unsafe.Sizeof(R.email))
}

func (R *Row) GetIdOffset() int {
	return 0
}

func (R *Row) GetUserNameOffset() int {
	return R.GetIdOffset() + R.GetIdSize()
}

func (R *Row) GetEmailOffset() int {
	return R.GetUserNameOffset() + R.GetUserNameSize()
}

func (R *Row) GetRowSize() int {
	return R.GetIdSize() + R.GetUserNameSize() + R.GetEmailSize()
}

func (R *Row) SerializeId(destination []byte) {
	var intByte []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(intByte, uint32(R.id))
	copy(destination, intByte)
}

func (R *Row) SerializeUserName(destination []byte) {
	copy(destination, R.userName[:])
}

func (R *Row) SerializeEmail(destination []byte) {
	copy(destination, R.email[:])
}

func (R *Row) DeSerializeId(destination []byte) {
	var id uint32
	id = binary.LittleEndian.Uint32(destination)
	R.id = int32(id)
}

func (R *Row) DeSerializeUserName(destination []byte) {
	copy(R.userName[:], destination)
}

func (R *Row) DeSerializeEmail(destination []byte) {
	copy(R.email[:], destination)
}
