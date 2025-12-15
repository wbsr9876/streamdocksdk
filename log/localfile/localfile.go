package localfile

import (
	"log"
	"os"
)

type Writer struct {
	path string
	file *os.File
}

func NewWriter(filePath string) *Writer {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Writer{
		path: filePath,
		file: file,
	}
}

func (w *Writer) LogMessage(message string) {
	_, _ = w.file.WriteString(message + "\n")
}

func (w *Writer) Close() {
	if w.file != nil {
		_ = w.file.Close()
	}
}
