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

func TestReadFirstLineOfSecondEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s7")

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
	assert.Equal(t, 2, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[1].data[0])
}

func TestReadSecondLineOfSecondEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s7")

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
	assert.Equal(t, 2, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[1].data[1])
}

func TestReadThirdLineOfSecondEntry(t *testing.T) {
	reader := NewFileDocumentReader("./s7")

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
	assert.Equal(t, 2, len(entries))
	assert.Equal(t, nil, err)
	assert.Equal(t, except, entries[1].data[2])
}

func TestReadMultipleEntriesInOneDocument(t *testing.T) {
	reader := NewFileDocumentReader("./s6")

	entries, err := reader.ReadEntries()

	assert.Equal(t, nil, err)
	assert.Equal(t, 12, len(entries))
	assert.Equal(t, "000000000", entries[0].toString())
	assert.Equal(t, "111111111", entries[1].toString())
	assert.Equal(t, "222222222", entries[2].toString())
	assert.Equal(t, "333333333", entries[3].toString())
	assert.Equal(t, "444444444", entries[4].toString())
	assert.Equal(t, "555555555", entries[5].toString())
	assert.Equal(t, "666666666", entries[6].toString())
	assert.Equal(t, "777777777", entries[7].toString())
	assert.Equal(t, "888888888", entries[8].toString())
	assert.Equal(t, "999999999", entries[9].toString())
	assert.Equal(t, "123456789", entries[10].toString())
	assert.Equal(t, "000000051", entries[11].toString())
}
