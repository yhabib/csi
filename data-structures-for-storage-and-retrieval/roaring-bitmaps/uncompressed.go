package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	return &uncompressedBitmap{}
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	idx := x / wordSize
	if idx >= uint32(len(b.data)) {
		return false
	}
	offset := x % wordSize
	return b.data[idx]&(1<<offset) != 0
}

func (b *uncompressedBitmap) Set(x uint32) {
	idx := x / wordSize
	offset := x % wordSize
	size := uint32(len(b.data))
	if idx > size {
		b.data = append(b.data, make([]uint64, 2*idx)...)
	}
	b.data[idx] |= (1 << offset)
}

// Go through both arrays and make usion of each bit
func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	if len(other.data) > len(b.data) {
		return other.Union(b)
	}

	var data []uint64
	data = append(data, b.data...)
	for i := 0; i < len(other.data); i++ {
		data[i] |= other.data[i]
	}

	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	if len(other.data) < len(b.data) {
		return other.Intersect(b)
	}
	var data []uint64
	data = append(data, b.data...)
	for i := 0; i < len(b.data); i++ {
		data[i] &= other.data[i]
	}
	return &uncompressedBitmap{
		data: data,
	}
}
