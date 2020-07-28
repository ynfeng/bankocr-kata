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
	fi, err := os.Open(reader.filePath)
	if err != nil {
		return nil, err
	}
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
		lineBytes := copyLineData(readBytes)
		entry.appendLineData(lineBytes)
		if entry.isDataComplete() {
			result = append(result, entry)
			entry = NewEmptyEntry()
		}
	}
	return result, nil
}

func copyLineData(readBytes []byte) [27]byte {
	var lineBytes [27]byte
	for i := 0; i < 27; i++ {
		lineBytes[i] = readBytes[i]
	}
	return lineBytes
}

func NewFileDocumentReader(filePath string) FileDocumentReader {
	return FileDocumentReader{filePath: filePath}
}
