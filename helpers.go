package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"log"
	"strings"
	"unicode/utf8"
)

func printable(s string) string {
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func uniqueSlice(arr []string) []string {
	mvalues := map[string]string{}
	for _, s := range arr {
		mvalues[s] = s
	}

	var filtered []string
	for _, s := range mvalues {
		filtered = append(filtered, s)
	}
	return filtered
}

//	func hashMD4(s string) string {
//		return fmt.Sprintf("%x", md4.Sum([]byte(s)))
//	}
func hashMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func hashSHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func hashSHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func hashSHA512(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

// SHA-224
func hashSHA224(s string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

// SHA-384
func hashSHA384(s string) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}

// SHA-512/224
func hashSHA512_224(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512_224([]byte(s)))
}

// SHA-512/256
func hashSHA512_256(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512_256([]byte(s)))
}

// SHA3-256
func hashSHA3_256(s string) string {
	hash := sha3.New256()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// SHA3-512
func hashSHA3_512(s string) string {
	hash := sha3.New512()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// RIPEMD-160
func hashRIPEMD160(s string) string {
	hash := ripemd160.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func permutations(input []string) [][]string {
	if len(input) > 4 {
		log.Fatal("Too many permutations")
	}
	var result [][]string
	heapPermutation(input, len(input), &result)
	return result
}

func heapPermutation(input []string, size int, result *[][]string) {
	// If size is 1, store the obtained permutation
	if size == 1 {
		tmp := make([]string, len(input))
		copy(tmp, input)
		*result = append(*result, tmp)
	}

	for i := 0; i < size; i++ {
		heapPermutation(input, size-1, result)

		// if size is odd, swap the first and last element
		// if size is even, swap the i-th and last element
		if size%2 == 1 {
			input[0], input[size-1] = input[size-1], input[0]
		} else {
			input[i], input[size-1] = input[size-1], input[i]
		}
	}
}

func combinations(arr []string) [][]string {
	var result [][]string

	n := len(arr)
	// Total combinations would be 2^n
	total := 1 << n

	// Iterate from 1 to 2^n
	for i := 1; i < total; i++ {
		var subset []string
		for j := 0; j < n; j++ {
			// If j-th bit in i is set, add arr[j] to subset
			if (i & (1 << j)) > 0 {
				subset = append(subset, arr[j])
			}
		}
		result = append(result, subset)
	}
	return result
}

func makeHashes(s string) []string {
	return []string{
		//hashMD4(s),
		hashMD5(s),
		hashSHA1(s),
		hashSHA256(s),
		hashSHA512(s),
		hashSHA224(s),
		hashSHA384(s),
		hashSHA3_256(s),
		hashSHA3_512(s),
		hashSHA512_224(s),
		hashSHA512_256(s),
	}
}
