package before

type location struct {
    current *piece
}

func NewLocation(piece *piece) *location {
    return &location{current: piece}
}