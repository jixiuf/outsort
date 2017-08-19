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
	bufSize     int
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
		bufSize:     bufSize,
		buffer:      make(IntSlice, 0, bufSize),
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
func (r *Reader) Reset() {
	r.buffer = r.buffer[0:0]
}
func (r *Reader) ReRead() error {
	r.buffer = r.buffer[0:0]
	return r.Read()
}

// 从文件中读取buffer对应长度的数据，放到buffer中，尽可能的填充满buffer(达到cap(buffer)后 就不到继续读取)
// 注意buffer中可能已经有部分内容，这部分内容会保留,
func (r *Reader) Read() (err error) {
	if r.fileHandler == nil {
		return errors.New("read error:no file handler")
	}
	if r.isEOF {
		// r.buffer = r.buffer[0:0] //
		return io.EOF
	}

	// r.buffer = r.buffer[0:cap(r.buffer)] //
	if cap(r.buffer) < 1 {
		return io.EOF
	}

	var scanSucc bool
	for { // read line by line
		if len(r.buffer) >= cap(r.buffer) {
			break
		}
		scanSucc = r.lineScanner.Scan()
		if !scanSucc {
			r.isEOF = true
			break
		}

		line := strings.TrimSpace(r.lineScanner.Text())
		if len(line) < 1 {
			continue
		}

		value, err := strconv.Atoi(line)
		if err == nil {
			r.buffer = append(r.buffer, int32(value))
		}
	}

	return nil
}
