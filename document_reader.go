package bankocr

type DocumentReader interface {
	ReadEntries() ([]Entry, error)
}
