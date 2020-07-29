package bankocr

import (
	"bufio"
	"fmt"
	"os"
)

type FileBankOutput struct {
	outputFile string
}

func (output FileBankOutput) appendLine(accountNo string) error {
	fi, _ := os.OpenFile(output.outputFile, os.O_WRONLY|os.O_APPEND, 0755)
	defer fi.Close()
	wr := bufio.NewWriter(fi)
	fmt.Fprintln(wr, accountNo)
	wr.Flush()
	return nil
}

func NewFileBankOutput(outputFile string) BankOutput {
	fi, _ := os.Create(outputFile)
	defer fi.Close()
	return FileBankOutput{outputFile: outputFile}
}
