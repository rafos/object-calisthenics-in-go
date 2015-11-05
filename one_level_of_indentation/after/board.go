package after

import "bytes"

type board struct {
	data [][]string
}

func NewBoard(data [][]string) *board {
	return &board{data: data}
}

func (b *board) Board() string {
	var buffer = &bytes.Buffer{}

	b.collectRows(buffer)

	return buffer.String()
}

func (b *board) collectRows(buffer *bytes.Buffer) {
	for i := 0; i < 10; i++ {
		b.collectRow(buffer, i)
	}
}

func (b *board) collectRow(buffer *bytes.Buffer, row int) {
	for j := 0; j < 10; j++ {
		buffer.WriteString(b.data[row][j])
	}

	buffer.WriteString("\n")
}
