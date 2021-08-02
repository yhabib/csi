package bitmap

const (
	wahWidth        = wordSize - 1
	msb      uint64 = 1 << wahWidth
	smsb     uint64 = msb >> 1
	allOn    uint64 = msb - 1
	allOff   uint64 = 0
)

// Precondition: runWord == allOn || runWord == allOff
func pack(runWord, runLength uint64) uint64 {
	if runWord == allOn {
		return msb | smsb | runLength
	} else {
		return msb | runLength
	}
}

func compress(b *uncompressedBitmap) []uint64 {
	// Step 1: Convert wordSize length words into wahWidth sized words
	var words []uint64
	offset := uint32(0)
	for offset < uint32(len(b.data))*wordSize {
		word := uint64(0)
		for i := 0; i < wahWidth; i++ {
			if b.Get(offset) {
				word |= 1 << i
			}
			offset++
		}
		words = append(words, word)
	}

	// Step 2: Handle consecutive runs of allOn and allOff words
	var compressed []uint64

	// Invariant: runWord == allOn || runWord == allOff
	var runWord, runLength uint64
	for _, word := range words {
		if word != runWord && runLength > 0 {
			compressed = append(compressed, pack(runWord, runLength))
			runLength = 0
		}
		if word == allOn || word == allOff {
			runWord = word
			runLength++
		} else {
			compressed = append(compressed, word)
		}
	}
	if runLength > 0 {
		compressed = append(compressed, pack(runWord, runLength))
	}
	return compressed
}

func isFill(word uint64) bool {
	return word&msb != 0
}

// Precondition: isFill(word)
func fillOn(word uint64) bool {
	return word&smsb != 0
}

// Precondition: isFill(word)
func fillRunLength(word uint64) uint64 {
	return word & ^(msb | smsb)
}

func decompress(compressed []uint64) *uncompressedBitmap {
	b := newUncompressedBitmap()
	offset := uint32(0)
	for _, word := range compressed {
		if isFill(word) {
			on := fillOn(word)
			runLength := fillRunLength(word)
			for i := uint64(0); i < runLength*wahWidth; i++ {
				if on {
					b.Set(offset)
				}
				offset++
			}
		} else { // literal
			for i := 0; i < wahWidth; i++ {
				if word&(1<<i) != 0 {
					b.Set(offset)
				}
				offset++
			}
		}
	}
	return b
}
