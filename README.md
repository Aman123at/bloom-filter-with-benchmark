# Bloom Filter in Golang

This documentation provides an overview of the **Bloom Filter** implementation in Golang. The Bloom Filter is a space-efficient probabilistic data structure used to test whether an element is part of a set. It can produce false positives but guarantees no false negatives. This implementation uses **SHA-1** for hashing.

## Key Components

### 1. `BloomFilter` Struct
The `BloomFilter` struct is the core data structure of the Bloom filter.

```go
type BloomFilter struct {
    bitset        []bool   // The bit array used to store element hashes
    size          uint     // The size of the bit array (m)
    hashFunctions []hash.Hash // Array of hash functions
    numHashes     int      // Number of hash functions (k)
}
```

### 2. `NewBloomFilter` Constructor
The `NewBloomFilter` function creates a new instance of a Bloom filter. It takes two parameters:
  - `n`: The expected number of elements to be inserted.
  - `p`: The desired false positive rate.

This function computes the optimal bit array size (`m`) and the number of hash functions (`k`) based on the inputs. It also initializes the bitset and generates the necessary hash functions.

```go
func NewBloomFilter(n uint, p float64) *BloomFilter
```
Example usage:

```go
n := uint(1000) // Expected number of elements
p := 0.05       // False positive rate
bloomFilter := NewBloomFilter(n, p)
```

### 3. `Add` Method
The `Add` method inserts an element into the Bloom filter. It hashes the data using multiple hash functions and sets the respective bits in the bit array.

```go
func (bf *BloomFilter) Add(data []byte)
```

Example usage:

```go
bloomFilter.Add([]byte("Apple"))
bloomFilter.Add([]byte("Cherry"))
```

### 4. `Check` Method
The `Check` method verifies whether an element might be in the Bloom filter. It hashes the data using the same hash functions and checks the bit array. If all corresponding bits are set, the element may be in the set; otherwise, it is definitely not in the set.

```go
func (bf *BloomFilter) Check(data []byte) bool
```

Example usage:

```go
result := bloomFilter.Check([]byte("Apple")) // true if present, false if definitely absent
```

### 5. `optimalSize` Function
The `optimalSize` function calculates the optimal size of the bit array (`m`) for a given number of elements (`n`) and the false positive rate (`p`).

```go
func optimalSize(n uint, p float64) uint
```

The formula used:

```go
m := -float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)
```

### 6. `optimalHashCount` Function
The `optimalHashCount` function calculates the optimal number of hash functions (`k`) based on the size of the bit array (`m`) and the number of elements (`n`).

```go
func optimalHashCount(m, n uint) int
```

The formula used:

```go
k := float64(m) / float64(n) * math.Ln2
```


# Benchmark 
`False Positive Rate` vs `Hash Functions`

<img width="574" alt="Screenshot 2024-09-16 at 6 00 19 PM" src="https://github.com/user-attachments/assets/ba76686b-a5f0-416e-a308-f9a1baf598ac">


# Conclusion

This implementation demonstrates a simple yet effective Bloom filter using **SHA-1** hash functions. The filter is space-efficient and designed for use cases where some false positives are acceptable but false negatives are not. You can adjust the size of the filter and the false positive rate by modifying the `n` and `p` parameters in the constructor.











