package after

import "bytes"

type location struct {
    current *piece
}

func NewLocation(piece *piece) *location {
    return &location{current: piece}
}

func (l *location) addTo(buffer *bytes.Buffer) {
    l.current.addTo(buffer)
}