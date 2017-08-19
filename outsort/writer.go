package outsort

import "os"
import "strconv"

type Writer struct {
	fileHandler *os.File
}

func NewWriter(fileName string) (w *Writer, err error) {
	fileHandler, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}

	w = &Writer{
		fileHandler: fileHandler,
	}

	return

}
func (w *Writer) Close() {
	if w.fileHandler != nil {
		w.fileHandler.Close()
		w.fileHandler = nil
	}
}

func (w *Writer) Write(data IntSlice) {
	for _, intValue := range data {
		w.fileHandler.WriteString(strconv.Itoa(int(intValue)))
		w.fileHandler.WriteString("\n")
	}
}
