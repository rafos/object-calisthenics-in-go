package before

type piece struct {
    representation string
}

func NewPiece(representation string) *piece {
    return &piece{representation: representation}
}