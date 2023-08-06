package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
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
		if s == "" {
			continue
		}
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
