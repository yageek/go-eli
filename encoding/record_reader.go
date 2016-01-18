package encoding

import "io"

const (
	maxPageBufferSize = 64
)

// Decoder helps reading record
type Decoder struct {
	r      io.Reader
	err    error
	buffer []byte
	pos    int
}

// NewReader creates a reader for decoding
// the different record
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

func (r *Decoder) Decode(v interface{}) error {
	packetType := make([]byte, 1)
	length := make([]byte, 1)

}
