package bankocr

type EntryDecorator interface {
	decorate(entry Entry) string
}

type DefaultEntryDecorator struct {
}

func NewDefaultEntryDecorator() EntryDecorator {
	return DefaultEntryDecorator{}
}

func (decorate DefaultEntryDecorator) decorate(entry Entry) string {
	if entry.isIllegal() {
		return entry.toString() + " ILL"
	}
	return entry.toString()
}
