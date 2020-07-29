package bankocr

type BankOutput interface {
	appendLine(accountNo string) error
}
