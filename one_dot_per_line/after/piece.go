package after
import "bytes"

type piece struct {
    representation string
}

func NewPiece(representation string) *piece {
    return &piece{representation: representation}
}

func (p *piece) character() string {
    return p.representation[0:1]
}

func (p *piece) addTo(buffer *bytes.Buffer) {
    buffer.WriteString(p.character())
}