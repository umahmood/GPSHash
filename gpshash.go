package gpshash

/*******************************************************************************

General purpose string hashing algorithms:

- http://www.partow.net/programming/hashfunctions/index.html#AvailableHashFunctions

Implemented in Go.

Note: The comments associated with each function are taken from:

- http://www.partow.net/programming/hashfunctions/index.html#AvailableHashFunctions

Full credit goes to Arash Partow.

*******************************************************************************/

// AP an algorithm produced by Arash Partow.
func AP(str string) uint {
	var hash uint = 0xAAAAAAAA
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint(str[i])*(hash>>3))
		} else {
			hash ^= (^((hash << 11) + (uint(str[i]) ^ (hash >> 5))))
		}
	}
	return hash
}

// BP a simple hash function taken from Bruno Preiss's book: 'Data Structures
// and Algorithms with Object-Oriented Design Patterns in Java,'' John Wiley &
// Sons, 2000.
func BP(str string) uint {
	var hash uint
	for i := 0; i < len(str); i++ {
		hash = hash<<7 ^ uint(str[i])
	}
	return hash
}

// BKDR this hash function comes from Brian Kernighan and Dennis Ritchie's book
// "The C Programming Language". It is a simple hash function using a strange
// set of possible seeds which all constitute a pattern of 31....31...31 etc, it
// seems to be very similar to the DJB hash function.
func BKDR(str string) uint {
	var seed uint = 131
	var hash uint

	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint(str[i])
	}
	return hash
}

// DJB produced by Professor Daniel J. Bernstein and shown first to the world
// on the usenet newsgroup comp.lang.c. It is one of the most efficient hash
// functions ever published.
func DJB(str string) uint {
	var hash uint = 5381
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint(str[i])
	}
	return hash
}

// DEK proposed by Donald E. Knuth in The Art Of Computer Programming Volume 3,
// under the topic of sorting and search chapter 6.4.
func DEK(str string) uint {
	var hash = uint(len(str))
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint(str[i])
	}
	return hash
}

// ELF similar to the PJW Hash function, but tweaked for 32-bit processors. It
// is a widely used hash function on UNIX based systems.
func ELF(str string) uint {
	var hash uint
	var x uint
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint(str[i])

		x = hash & 0xF0000000

		if x != 0 {
			hash ^= (x >> 24)
		}
		hash &= ^x
	}
	return hash
}

// FNV Fowler-Noll-Vo is a non-cryptographic hash function created by Glenn
// Fowler, Landon Curt Noll, and Kiem-Phong Vo.
func FNV(str string) uint {
	const fnvPrime = uint(0x811C9DC5)
	var hash uint
	for i := 0; i < len(str); i++ {
		hash *= fnvPrime
		hash ^= uint(str[i])
	}
	return hash
}

// JS a bitwise hash function created by Justin Sobel
func JS(str string) uint {
	var hash uint = 1315423911
	for i := 0; i < len(str); i++ {
		hash ^= ((hash << 5) + uint(str[i]) + (hash >> 2))
	}
	return hash
}

// PJW based on work by Peter J. Weinberger of AT&T Bell Labs. The book
// Compilers (Principles, Techniques and Tools) by Aho, Sethi and Ulman,
// recommends the use of hash functions that employ the hashing methodology
// found in this particular algorithm.
func PJW(str string) uint {
	const maxUint = ^uint(0)
	const bitsPerWord = uint(32 * (1 + maxUint>>63))
	const threeQuarters = uint((bitsPerWord * 3) / 4)
	const oneEighth = uint(bitsPerWord / 8)
	const highBits = uint(1<<(bitsPerWord-oneEighth) - 1)

	var hash uint
	var test uint
	for i := 0; i < len(str); i++ {
		hash = (hash << oneEighth) + uint(str[i])
		test = hash & highBits
		if test != 0 {
			hash ^= (test >> threeQuarters) & (^highBits)
		}
	}
	return hash
}

// RS a simple hash function from Robert Sedgwicks Algorithms in C book.
func RS(str string) uint {
	var b uint = 378551
	var a uint = 63689
	var hash uint
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint(str[i])
		a = a * b
	}
	return hash
}

// SDBM is the algorithm of choice which is used in the open source SDBM project.
// The hash function seems to have a good over-all distribution for many
// different data sets.
func SDBM(str string) uint {
	var hash uint
	for i := 0; i < len(str); i++ {
		hash = uint(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
