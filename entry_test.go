package bankocr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryToAccountString(t *testing.T) {
	data := [3][27]byte{
		{' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|',},
		{' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "333333333", entry.toString())
}

func TestGetNumberData(t *testing.T) {
	data := [3][27]byte{
		{' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|',},
		{' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	except := [3][3]byte{
		{' ', '_', ' ',},
		{' ', '_', '|',},
		{' ', '_', '|',},
	}
	assert.Equal(t, except, entry.getNumberData(1))
	assert.Equal(t, except, entry.getNumberData(2))
	assert.Equal(t, except, entry.getNumberData(3))
	assert.Equal(t, except, entry.getNumberData(4))
	assert.Equal(t, except, entry.getNumberData(5))
	assert.Equal(t, except, entry.getNumberData(6))
	assert.Equal(t, except, entry.getNumberData(7))
	assert.Equal(t, except, entry.getNumberData(8))
	assert.Equal(t, except, entry.getNumberData(9))
}

func TestGetNumber1(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "1", entry.GetNumber(1))
}

func TestGetNumber2(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "2", entry.GetNumber(2))
}

func TestGetNumber3(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "3", entry.GetNumber(3))
}

func TestGetNumber4(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "4", entry.GetNumber(4))
}

func TestGetNumber5(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "5", entry.GetNumber(5))
}

func TestGetNumber6(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "6", entry.GetNumber(6))
}

func TestGetNumber7(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "7", entry.GetNumber(7))
}

func TestGetNumber8(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "8", entry.GetNumber(8))
}

func TestGetNumber9(t *testing.T) {
	data := [3][27]byte{
		{' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', ' ', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{' ', ' ', '|', ' ', '_', '|', ' ', '_', '|', '|', '_', '|', '|', '_', ' ', '|', '_', ' ', ' ', ' ', '|', '|', '_', '|', '|', '_', '|',},
		{' ', ' ', '|', '|', '_', ' ', ' ', '_', '|', ' ', ' ', '|', ' ', '_', '|', '|', '_', '|', ' ', ' ', '|', '|', '_', '|', ' ', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "8", entry.GetNumber(8))
}

func TestGetNumber0(t *testing.T) {
	data := [3][27]byte{
		{' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',},
		{'|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|',},
		{'|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|',},
	}
	entry := NewEntry(data)

	assert.Equal(t, "0", entry.GetNumber(3))
}

func TestAppendLineData(t *testing.T) {
	lineData1 := [27]byte {' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ', ' ', '_', ' ',}
	lineData2 := [27]byte {'|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|', '|', ' ', '|',}
	lineData3 := [27]byte {'|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|', '|', '_', '|',}

	entry := NewEmptyEntry()
	
	entry.appendLineData(lineData1)
	entry.appendLineData(lineData2)
	entry.appendLineData(lineData3)


	assert.Equal(t, "0", entry.GetNumber(3))
	assert.Equal(t, true, entry.isDataComplete())
}