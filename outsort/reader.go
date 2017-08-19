package outsort

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// 从文件读取指定长度的数组放到 reader.buffer中
type Reader struct {
	buffer      IntSlice
	fileHandler *os.File
	lineScanner *bufio.Scanner // reader line by line
	pos         int
	isEOF       bool // 是否已经读取到文件末尾
}

func NewReader(fileName string, bufSize int) (reader *Reader, err error) {
	fileHandler, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader = &Reader{
		buffer:      make(IntSlice, bufSize, bufSize),
		fileHandler: fileHandler,
		lineScanner: bufio.NewScanner(fileHandler),
		pos:         0,
		isEOF:       false,
	}

	return reader, nil
}
func (r *Reader) Close() {
	if r.fileHandler != nil {
		r.fileHandler.Close()
		r.fileHandler = nil
		r.lineScanner = nil
	}
}

func (r *Reader) Data() (arr IntSlice) {
	return r.buffer
}
func (r *Reader) IsEOF() bool {
	return r.isEOF
}

// 从文件中读取buffer对应长度的数据，放到buffer中，尽可能的填充满buffer
func (r *Reader) Read() (err error) {
	if r.fileHandler == nil {
		return errors.New("read error:no file handler")
	}
	if r.isEOF {
		r.buffer = r.buffer[0:0] //
		return io.EOF
	}

	r.buffer = r.buffer[0:cap(r.buffer)] //
	var idx int
	if len(r.buffer) < 1 {
		return io.EOF
	}

	var scanSucc bool
	for { // read line by line
		scanSucc = r.lineScanner.Scan()
		if !scanSucc {
			break
		}

		line := strings.TrimSpace(r.lineScanner.Text())
		if len(line) < 1 {
			continue
		}

		value, err := strconv.Atoi(line)
		if err == nil {
			r.buffer[idx] = int32(value)
			idx++
		}
		if idx >= cap(r.buffer) {
			break
		}
	}
	if !scanSucc {
		r.isEOF = true
	}

	r.buffer = r.buffer[0:idx]
	return r.lineScanner.Err()
}
