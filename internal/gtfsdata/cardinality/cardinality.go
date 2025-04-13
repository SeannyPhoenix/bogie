package cardinality

import (
	"encoding/binary"
	"math"
	"math/bits"
)

type ChunkSize int

const (
	Byte ChunkSize = 1
	KB   ChunkSize = Byte * 1024
	MB   ChunkSize = KB * 1024
	GB   ChunkSize = MB * 1024
	TB   ChunkSize = GB * 1024
)

type Cardinality struct {
	Buckets int
	Bits    int
}

func GetCardinality(chunkSize ChunkSize, dataSize int, partitionCount int) Cardinality {
	var c Cardinality

	if dataSize == 0 || chunkSize == 0 || partitionCount == 0 {
		return c
	}

	chunks := math.Ceil(float64(dataSize) / float64(chunkSize))
	min := math.Min(chunks, float64(partitionCount))
	n := math.Ceil(min)

	c.Bits = bits.Len(uint(n - 1))
	c.Buckets = 1 << c.Bits

	return c
}

func GetPrefix(id []byte, bits int) uint64 {
	return binary.BigEndian.Uint64(id[:8]) >> (64 - bits)
}
