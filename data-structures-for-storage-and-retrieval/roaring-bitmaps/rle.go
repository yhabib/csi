package bitmap

const (
	wahWidth uint64 = wordSize - 1
	msb      uint64 = 1 << wahWidth
	smsb     uint64 = msb >> 1
	allOn    uint64 = msb - 1
	allOff   uint64 = 0
)

// it is either allOn or allOff
func pack(word uint64, length uint64) uint64 {
	// mask := uint64(1 << (wahWidth - 1))
	if word == allOn {
		// set the two highest bits to 11 and then set number of matches
		// mask |= 1 << (wahWidth - 2)
		// return mask | length
		return msb | smsb | length
	}
	// mask = mask >> 1
	// // set highest bit to 1 next to 0 and then set number of matches
	// return mask | length
	return msb | length
}

// 1. Convert slice into wah words
// 2. Compress based on values
func compress(b *uncompressedBitmap) []uint64 {
	var wahWords []uint64
	var offset uint32

	// 1.
	for offset < uint32(len(b.data)*wordSize) {
		word := uint64(0)
		for i := uint64(0); i < wahWidth; i++ {
			if b.Get(offset) {
				word |= 1 << i
			}
			offset++
		}
		wahWords = append(wahWords, word)
	}

	// 2.
	var compressed []uint64
	var currentWord, length uint64
	for _, word := range wahWords {
		if currentWord != word && length > 0 {
			compressed = append(compressed, pack(currentWord, length))
			length = 0
		}
		if word == allOff || word == allOn {
			currentWord = word
			length++
		} else {
			compressed = append(compressed, word)
		}
	}
	if length > 0 {
		compressed = append(compressed, pack(currentWord, length))
	}

	return compressed
}

func getValueAndReps(word uint64) (bool, uint64) {
	isOn := word&smsb != 0
	numOfRepetitions := word & (allOn >> 2)

	return isOn, numOfRepetitions
}

func decompress(compressed []uint64) *uncompressedBitmap {
	uncompressedBitmap := newUncompressedBitmap()
	var offset uint32
	for _, wahWord := range compressed {
		if (wahWord & msb) != 0 {
			isOn, numOfRepetitions := getValueAndReps(wahWord)
			for i := 0; i < int(numOfRepetitions*wahWidth); i++ {
				if isOn {
					uncompressedBitmap.Set(offset)
				}
				offset++
			}
		} else {
			for i := 0; i < int(wahWidth); i++ {
				if wahWord&(1<<i) != 0 {
					uncompressedBitmap.Set(offset)
				}
				offset++
			}
		}
	}

	return uncompressedBitmap
}
