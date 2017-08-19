package outsort

import (
	"errors"
	"io"
	"os"
	"strconv"
)

type Sort struct {
	bufferSize   int
	tmpFileNames []string
	inFileName   string
	outFileName  string
	tmpFileIdx   int
	resultBuffer IntSlice
}

const (
	sizeOfInt = 4
)

func NewSort(bufferSize int) (s *Sort) {
	s = &Sort{
		bufferSize: bufferSize,
	}
	if s.bufferSize < 2 {
		s.bufferSize = 2
	}
	s.resultBuffer = make(IntSlice, 0, s.bufferSize*2)

	return
}
func (s *Sort) Sort(inFileName, outFileName string) error { // 传入待排序的文件(文件每行一个int数字)，及排序后存放结果的文件
	s.inFileName = inFileName
	s.outFileName = outFileName
	err := s.initSort()
	if err != nil {
		return err
	}

	return s.mergeAll()
}

func (s *Sort) initSort() (err error) { //
	// 从待排序文件中分段读取数据，分段排序后 生成临时的已排序的小文件
	// 并记录这些临时文件的名称，以便进一步排序

	reader, err := NewReader(s.inFileName, s.bufferSize)
	if err != nil {
		return err
	}
	defer reader.Close()

	for {
		err := reader.ReRead()
		if err != nil && err != io.EOF { // 出错
			return err
		}
		arr := reader.Data()
		if len(arr) == 0 {
			break
		}
		tmpFileName := s.getTmpFileName()
		tmpFileWriter, err := NewWriter(tmpFileName)
		if err != nil {
			return err
		}

		// 内存排序，并写到临时文件中
		arr.Sort()
		tmpFileWriter.Write(arr)
		tmpFileWriter.Close()
		s.tmpFileNames = append(s.tmpFileNames, tmpFileName)
		if reader.IsEOF() {
			break
		}

	}
	return nil

}
func (s *Sort) mergeAll() (err error) { //
	if len(s.tmpFileNames) == 0 {
		return errors.New("no data to merge and sort")
	}
	if len(s.tmpFileNames) == 1 {
		os.Rename(s.tmpFileNames[0], s.outFileName)
	}
	var leftFileName string = s.tmpFileNames[0]
	for _, rightFileName := range s.tmpFileNames[1:] {
		newTmpFile, err := s.mergeTwoFiles(leftFileName, rightFileName)
		if err != nil || newTmpFile == "" {
			continue
		}
		leftFileName = newTmpFile
	}
	os.Rename(leftFileName, s.outFileName)

	return nil

}
func (s *Sort) mergeTwoFiles(tmpFileName1, tmpFileName2 string) (newTmpFileName string, err error) { //
	leftReader, err := NewReader(tmpFileName1, s.bufferSize)
	if err != nil {
		return "", err
	}
	defer leftReader.Close()

	rightReader, err := NewReader(tmpFileName2, s.bufferSize)
	if err != nil {
		return "", err
	}
	defer rightReader.Close()

	newTmpFileName = s.getTmpFileName()
	tmpFileWriter, err := NewWriter(newTmpFileName)
	if err != nil {
		return newTmpFileName, err
	}
	defer tmpFileWriter.Close()
	s.merge(leftReader, rightReader, tmpFileWriter)
	// merge之后，删除tmpFileName1  tmpFileName2 两个临时文件，不再有用
	os.Remove(tmpFileName1)
	os.Remove(tmpFileName2)

	return

}
func (s *Sort) merge(leftReader, rightReader *Reader, tmpFileWriter *Writer) { //
	for {
		s.resultBuffer = s.resultBuffer[0:0] // 将buffer缩为长度0
		if leftReader.IsEOF() && rightReader.IsEOF() {
			break // 已有文件读到末尾
		}

		leftReader.Read()  // 从文件读取内容填充满 leftReader.buffer
		rightReader.Read() // 从文件读取内容填充满 rightReader.buffer

		Merge(&(leftReader.buffer), &(rightReader.buffer), &(s.resultBuffer)) // 内存内归并两个数组
		tmpFileWriter.Write(s.resultBuffer)                                   // 将归并后的数组写回到文件
		if len(s.resultBuffer) == 0 {                                         // 合并的结果为空，说明至少有一个队列已经掏空，此时只需要把另一个队列的内容append 到最后即可
			if leftReader.IsEOF() && len(leftReader.buffer) == 0 {
				tmpFileWriter.Write(rightReader.buffer) //
				rightReader.Reset()
			}
			if rightReader.IsEOF() && len(rightReader.buffer) == 0 {
				tmpFileWriter.Write(leftReader.buffer) //
				leftReader.Reset()
			}
		}
	}

	for !leftReader.IsEOF() || len(leftReader.buffer) != 0 {
		leftReader.Read()
		tmpFileWriter.Write(leftReader.buffer)
		leftReader.Reset()
	}
	for !rightReader.IsEOF() || len(rightReader.buffer) != 0 {
		rightReader.Read()
		tmpFileWriter.Write(rightReader.buffer)
		rightReader.Reset()
	}

	return
}
func (s *Sort) getTmpFileName() string { //
	s.tmpFileIdx++
	return "tmp_" + strconv.Itoa(s.tmpFileIdx)

}
