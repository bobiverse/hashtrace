package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
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

// func hashMD4(s string) string {
// 	return fmt.Sprintf("%x", md4.Sum([]byte(s)))
// }
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
