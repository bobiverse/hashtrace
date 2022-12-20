package main

import (
	"crypto/md5"
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

func hashMD5(s string) string {
	buf := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", buf)
}
