package fenwick

import "errors"

var IndexError = errors.New("Index is out of range")

type Fenwick struct {
	data []int
}

func New(size uint) *Fenwick {
	var f Fenwick
	f.data = make([]int, size)
	return &f
}

func (f *Fenwick) Add(pos int, value int) error {
	if pos < 0 || pos >= len(f.data) {
		return IndexError
	}
	for ; pos < len(f.data); pos = pos | (pos + 1) {
		f.data[pos] += value
	}
	return nil
}

func (f *Fenwick) Sum(begin, end int) int {
	endSum := f.prefixSum(end)
	beginSum := f.prefixSum(begin - 1)
	return endSum - beginSum
}

func (f *Fenwick) prefixSum(pos int) (result int) {
	if pos >= len(f.data) {
		pos = len(f.data) - 1
	}
	for ; pos >= 0; pos = (pos & (pos + 1)) - 1 {
		result += f.data[pos]
	}
	return
}
