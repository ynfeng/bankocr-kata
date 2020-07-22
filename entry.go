package bankocr

import (
	"strconv"
)

type Entry struct {
	data [3][27]byte
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

func NewEntry(data [3][27]byte) Entry {
	return Entry{data: data}
}
