package bankocr

type BankOcr struct {
	inputReader FileDocumentReader
}

func (ocr BankOcr) getAccountNo(line int) (string, error) {
	entries, _ := ocr.inputReader.ReadEntries()
	return entries[line-1].toString(), nil
}

func (ocr BankOcr) output(output BankOutput, decorator EntryDecorator) error {
	entries, _ := ocr.inputReader.ReadEntries()
	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		output.appendLine(decorator.decorate(entry))
	}
	return nil
}

func NewBankOcr(reader FileDocumentReader) BankOcr {
	return BankOcr{inputReader: reader}
}
