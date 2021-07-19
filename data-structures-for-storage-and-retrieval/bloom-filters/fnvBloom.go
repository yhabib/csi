package main

import (
	"encoding/binary"
	"hash/fnv"
	"unsafe"
)

const (
	BYTES_PER_ELEMENT = 8
	BITS_PER_ELEMENT  = 64
)

type fnvBloomFilter struct {
	data   []uint64
	nHash  uint8
	length uint32
}

func newFnvBloomFilter(size int, nHash uint8) *fnvBloomFilter {
	numElements := size / BYTES_PER_ELEMENT
	length := uint32(numElements * BITS_PER_ELEMENT)
	return &fnvBloomFilter{
		data:   make([]uint64, numElements),
		nHash:  nHash,
		length: length,
	}
}

// There is no seed for fnv to "create" multiple hash functions with same one
// Based on solution :(
func (b *fnvBloomFilter) hash(item string, i int) uint32 {
	fnv := fnv.New32()

	// this change allows to avoid copying the data and just reference the underlying bits
	// bytes := []byte(item)
	bytes := *((*[]byte)(unsafe.Pointer(&item)))
	prefix := byte(i)
	fnv.Write([]byte{prefix})
	fnv.Write(bytes)
	return fnv.Sum32() % b.length
}

func (b *fnvBloomFilter) get(k uint32) bool {
	index := k / BITS_PER_ELEMENT
	offset := k % BITS_PER_ELEMENT
	return b.data[index]&(1<<offset) != 0
}

func (b *fnvBloomFilter) set(k uint32) {
	index := k / BITS_PER_ELEMENT
	offset := k % BITS_PER_ELEMENT
	b.data[index] = b.data[index] | (1 << offset)
}

func (b *fnvBloomFilter) add(item string) {
	for i := 0; i < int(b.nHash); i++ {
		idx := b.hash(item, i)
		b.set(idx)
	}
}

func (b *fnvBloomFilter) maybeContains(item string) bool {
	for i := 0; i < int(b.nHash); i++ {
		idx := b.hash(item, i)
		if !b.get(idx) {
			return false
		}
	}
	return true
}

func (b *fnvBloomFilter) memoryUsage() int {
	return binary.Size(b.data)
}
