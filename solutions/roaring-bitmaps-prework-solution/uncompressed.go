package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	return &uncompressedBitmap{}
}

func indexAndShift(x uint32) (uint32, uint32) {
	return x / wordSize, x % wordSize
}

func nextPowOf2(n uint32) uint32 {
	result := uint32(1)
	for n >= result {
		result <<= 1
	}
	return result
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	index, shift := indexAndShift(x)
	if index >= uint32(len(b.data)) {
		return false
	}
	return b.data[index]&(1<<shift) != 0
}

func (b *uncompressedBitmap) Set(x uint32) {
	index, shift := indexAndShift(x)
	n := uint32(len(b.data))
	if index >= n {
		b.data = append(b.data, make([]uint64, nextPowOf2(index)-n)...)
	}
	b.data[index] |= 1 << shift
}

func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	// Ensure that len(other.data) <= len(b.data)
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
	// Ensure that len(other.data) <= len(b.data)
	if len(b.data) < len(other.data) {
		return other.Intersect(b)
	}

	var data []uint64
	data = append(data, other.data...)
	for i := 0; i < len(other.data); i++ {
		data[i] &= b.data[i]
	}
	return &uncompressedBitmap{
		data: data,
	}
}
