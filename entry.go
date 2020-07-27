package bankocr

import (
	"strconv"
)

type Entry struct {
	data       [3][27]byte
	lineDataNo int
}

func (entry *Entry) getNumberData(pos int) [3][3]byte {
	var result [3][3]byte
	data := entry.data
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = data[i][(pos-1)*3+j]
		}
	}
	return result
}

func (entry *Entry) GetNumber(pos int) string {
	data := entry.getNumberData(pos)
	for n, numData := range NUMBERS {
		find := true
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if data[i][j] != numData[i][j] {
					find = false
					break
				}
			}
			if find == false {
				break
			}
		}
		if find {
			return strconv.Itoa(n)
		}
	}
	return ""
}

func (entry *Entry) toString() string {
	result := ""
	for i := 1; i <= 9; i++ {
		number := entry.GetNumber(i)
		result += number
	}
	return result
}

func (entry *Entry) appendLineData(lineData [27]byte) {
	entry.data[entry.lineDataNo] = lineData
	entry.lineDataNo++
}

func (entry *Entry) isDataComplete() bool {
	if entry.lineDataNo == 3 {
		return true
	}
	return false
}

func NewEntry(data [3][27]byte) Entry {
	return Entry{data: data, lineDataNo: 3}
}

func NewEmptyEntry() Entry {
	var lineData [3][27]byte
	return Entry{data: lineData}
}
