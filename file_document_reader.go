package bankocr

import (
	"io/ioutil"
)

type FileDocumentReader struct {
	filePath string
}

func (reader FileDocumentReader) ReadEntries() ([]Entry, error) {
	result := make([]Entry, 0)

	bytes, _ := ioutil.ReadFile(reader.filePath)
	totalBytes := len(bytes)
	readIdx := 0
	lineNo := 0
	for readIdx < totalBytes {
		var readBytes [3][27]byte
		for i := 0; i < 3; i++ {
			col := 0
			for {
				readChar := bytes[28*i+col+(lineNo*84)]
				if readChar == 0xa {
					break
				}
				readBytes[i][col] = readChar
				col++
			}
		}
		result = append(result, NewEntry(readBytes))
		readIdx += 85
		lineNo++
	}
	return result, nil
}

func NewFileDocumentReader(filePath string) FileDocumentReader {
	return FileDocumentReader{filePath: filePath}
}
