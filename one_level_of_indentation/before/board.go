package before

import "bytes"

type board struct {
	data [][]string
}

func NewBoard(data [][]string) *board {
	return &board{data: data}
}

func (b *board) Board() string {
	var buffer = &bytes.Buffer{}

	// 0
	for i := 0; i < 10; i++ {
		// 1
		for j := 0; j < 10; j++ {
			// 2
			buffer.WriteString(b.data[i][j])
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}
