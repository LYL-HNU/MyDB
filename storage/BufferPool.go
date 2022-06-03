package storage

import (
	"fmt"
	"os"
)

type BufferPool struct {
	fileDesc   *os.File
	fileLength uint32
	pages      [TableMaxPages]*Page
}

func (bufferPool *BufferPool) GetFileLength() uint32 {
	return bufferPool.fileLength
}

func (bufferPool *BufferPool) SetFileDesc(file *os.File) {
	bufferPool.fileDesc = file
}

func (bufferPool *BufferPool) SetFileLength(length uint32) {
	bufferPool.fileLength = length
}

func (bufferPool *BufferPool) InitBufferPool() {
	for i := 0; i < TableMaxPages; i++ {
		bufferPool.pages[i] = nil
	}
}

func (bufferPool *BufferPool) GetPage(pageNum uint32) *Page {
	if pageNum >= TableMaxPages {
		fmt.Printf("Tried to fetch page number out of bounds. %d > %d\\n", pageNum, TableMaxPages)
		os.Exit(2)
	}
	//对应页未加载到内存
	if bufferPool.pages[pageNum] == nil {
		numPages := bufferPool.fileLength / PageSize
		//尾部存在部分页
		if bufferPool.fileLength%PageSize != 0 {
			numPages += 1
		}
		page := new(Page)
		pageByte := make([]byte, PageSize)
		L, err := bufferPool.fileDesc.ReadAt(pageByte, int64(pageNum*PageSize))
		if err != nil {
			fmt.Printf("getPage Error: %s \n", err.Error())
			os.Exit(3)
		}
		fmt.Printf("read Page Size is: %d \n", L)
		page.SetPage(pageByte)
		bufferPool.pages[pageNum] = page
	}
	return bufferPool.pages[pageNum]
}

func (bufferPool *BufferPool) FlushBack(pageNum uint32) {
	pageByte := make([]byte, PageSize)
	tmp := bufferPool.pages[pageNum].Page
	for i := 0; i < len(tmp); i++ {
		pageByte[i] = tmp[i]
	}
	_, err := bufferPool.fileDesc.Write(pageByte)
	if err != nil {
		fmt.Printf("Wrong Flush Back")
	}
}
