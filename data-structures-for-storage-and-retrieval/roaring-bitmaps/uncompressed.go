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

func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}
