package bankocr

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBankOcr(t *testing.T) {
	reader := NewFileDocumentReader("./s6")
	ocr := NewBankOcr(reader)

	accountNo, err := ocr.getAccountNo(12)

	assert.Equal(t, nil, err)
	assert.Equal(t, "000000051", accountNo)
}

func TestWriteToFile(t *testing.T) {
	output := NewFileBankOutput("./output")
	reader := NewFileDocumentReader("./s6")

	ocr := NewBankOcr(reader)
	err := ocr.output(output, NewDefaultEntryDecorator())

	accountNo := getFromFile("./output", 12)

	assert.Equal(t, nil, err)
	assert.Equal(t, "000000051", accountNo)
}

func TestWriteToFileWithIllegalAccountNo(t *testing.T) {
	output := NewFileBankOutput("./output")
	reader := NewFileDocumentReader("./s8")

	ocr := NewBankOcr(reader)
	err := ocr.output(output, NewDefaultEntryDecorator())

	accountNo := getFromFile("./output", 1)

	assert.Equal(t, nil, err)
	assert.Equal(t, "1234?678? ILL", accountNo)
}

func getFromFile(file string, line int) string {
	result := ""
	fi, _ := os.Open(file)
	defer fi.Close()
	br := bufio.NewReader(fi)
	currentLine := 0
	for {
		readBytes, _, _ := br.ReadLine()
		currentLine++
		if currentLine == line {
			return string(readBytes)
		}
	}
	return result
}
