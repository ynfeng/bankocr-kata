package bankocr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFirstLineOfEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s1")

	entries, err := reader.ReadEntries()

	except := [27]byte{
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' ',
		' ', '_', ' '}
	assert.Equal(t, 1, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[0].data[0])
}

func TestReadSecondLineOfEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s1")

	entries, err := reader.ReadEntries()

	except := [27]byte{
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
	}
	assert.Equal(t, 1, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[0].data[1])
}

func TestReadThirdLineOfEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s1")

	entries, err := reader.ReadEntries()

	except := [27]byte{
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
		' ', '_', '|',
	}
	assert.Equal(t, 1, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[0].data[2])
}

func TestReadMultipleEntriesInOneDocument(t *testing.T) {
	reader := NewFileDocumentReader("./s3")

	entries, err := reader.ReadEntries()

	assert.Equal(t, nil, err)
	assert.Equal(t, 3, len(entries))
}
