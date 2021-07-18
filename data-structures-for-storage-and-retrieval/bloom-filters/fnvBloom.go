package main

import (
	"encoding/binary"
	"hash/fnv"
)

type fnvBloomFilter struct {
	data  []bool
	nHash uint8
	size  uint32
}

func newFnvBloomFilter(size uint32, nHash uint8) *fnvBloomFilter {
	return &fnvBloomFilter{
		data:  make([]bool, size),
		nHash: nHash,
		size:  size,
	}
}

// There is no seed for fnv to "create" multiple hash functions with same one
// Based on solution :(
func (b *fnvBloomFilter) hash(item string, i int) uint32 {
	fnv := fnv.New32()
	bytes := []byte(item)
	prefix := byte(i)
	fnv.Write([]byte{prefix})
	fnv.Write(bytes)
	return fnv.Sum32() % uint32(b.size)
}

func (b *fnvBloomFilter) get(index uint32) bool {
	return b.data[index]
}

func (b *fnvBloomFilter) set(index uint32) {
	b.data[index] = true
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
