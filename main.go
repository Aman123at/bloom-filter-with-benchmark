package main

import (
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"log"
	"math"
)

type BloomFilter struct {
	bitset        []bool
	size          uint
	hashFunctions []hash.Hash
	numHashes     int
}

func NewBloomFilter(n uint, p float64) *BloomFilter {
	m := optimalSize(n, p)      // optimal size of bitarray
	k := optimalHashCount(m, n) // number of hash functions

	bloom := &BloomFilter{
		bitset:    make([]bool, m),
		size:      m,
		numHashes: k,
	}

	// Generate multiple hash functions (for simplicity, we use SHA-1 and different prefixes for each hash function)
	for i := 0; i < k; i++ {
		bloom.hashFunctions = append(bloom.hashFunctions, sha1.New())
	}

	return bloom
}

func (bf *BloomFilter) hash(i int, data []byte) uint64 {
	// Copy the hash function and prefix the data with the hash function index
	h := bf.hashFunctions[i]
	h.Reset()
	h.Write([]byte{byte(i)})
	h.Write(data)
	hashSum := h.Sum(nil)
	// Convert the hash sum to an unsigned integer
	return binary.BigEndian.Uint64(hashSum)
}

func (bf *BloomFilter) Add(data []byte) {
	for i := 0; i < bf.numHashes; i++ {
		hashValue := bf.hash(i, data)
		position := hashValue % uint64(bf.size)
		bf.bitset[position] = true
	}
}

func (bf *BloomFilter) Check(data []byte) bool {
	for i := 0; i < bf.numHashes; i++ {
		hashValue := bf.hash(i, data)
		position := hashValue % uint64(bf.size)
		if !bf.bitset[position] {
			return false // if any position is not set, the element is definitely not in the set
		}
	}
	return true // element might be in the set
}

func optimalSize(n uint, p float64) uint {
	m := -float64(n) * math.Log(p) / math.Ln2 * math.Ln2 // math.Ln2 ~ 0.63
	return uint(math.Ceil(m))
}

func optimalHashCount(m, n uint) int {
	k := float64(m) / float64(n) * math.Ln2
	return int(math.Ceil(k))
}

func main() {
	n := uint(1000) // number of expected elements
	p := 0.05       // false positive probability
	bFilter := NewBloomFilter(n, p)
	log.Println("bf-size", bFilter.size)
	log.Println("bf-nofhashes", bFilter.numHashes)
	log.Println("bf-hashfuncs", bFilter.hashFunctions)
	log.Println("bf-bitset", len(bFilter.bitset))
	bFilter.Add([]byte("Apple"))
	bFilter.Add([]byte("Cherry"))
	bFilter.Add([]byte("Peach"))

	log.Println(bFilter.Check([]byte("Apple")))  // true
	log.Println(bFilter.Check([]byte("Banana"))) // expected false but might be true based on n and p config
}
