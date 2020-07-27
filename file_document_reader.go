package bankocr

import (
	"bufio"
	"io"
	"os"
)

type FileDocumentReader struct {
	filePath string
}

func (reader FileDocumentReader) ReadEntries() ([]Entry, error) {
	result := make([]Entry, 0)
	fi, _ := os.Open(reader.filePath)
	defer fi.Close()

	br := bufio.NewReader(fi)
	entry := NewEmptyEntry()
	for {
		readBytes, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(readBytes) == 0 {
			continue
		}
		var lineBytes [27]byte
		for i := 0; i < 27; i++ {
			lineBytes[i] = readBytes[i]
		}
		entry.appendLineData(lineBytes)
		if entry.isDataComplete() {
			result = append(result, entry)
			entry = NewEmptyEntry()
		}
	}
	return result, nil
}

func NewFileDocumentReader(filePath string) FileDocumentReader {
	return FileDocumentReader{filePath: filePath}
}
