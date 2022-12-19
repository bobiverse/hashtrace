package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var hash string
	var data string
	flag.StringVar(&hash, "hash", "", "target hash to crack")
	flag.StringVar(&data, "data", "", "data possibly used to hash")
	flag.Parse()

	if hash == "" || data == "" {
		log.Fatalf("Usage:\n\t./hashguess -hash='90b76b4e' -data='firetruck'\n")
	}

	fmt.Println(">> " + hash)
	fmt.Println(">> [" + data + "]")
	log.Println(strings.Repeat(".", 80))

	// generate needles
	var needles []string
	needles = append(needles, data)
	needles = append(needles, reverse(data))
	needles = append(needles, splitToNeedles(data)...)
	needles = append(needles, splitToNeedles(reverse(data))...)

	needles = append(needles, addSuffixes(needles, []string{"\n", "\r", "\r\n", "\n\r"})...)

	needles = uniqueSlice(needles)

	N := len(hash)
	for _, s := range needles {

		// hashesh to check
		hashes := []string{
			hashMD5(s),
		}

		// if source hash is shorter
		for _, checkHash := range hashes {
			check(s, hash, checkHash)

			if N >= len(checkHash) {
				continue
			}

			n := len(checkHash)
			for i := 0; i < N; i++ {
				if i > n {
					continue
				}
				partHash := checkHash[i : i+N]
				if strings.Contains(hash, partHash) {
					check(s, hash, partHash)
				}
			}

		}
	}

	log.Fatalf("Done. NOT FOUND.")

}

// abcdefg => [a,b,c,d,e,f,g, ab, bc, .., abc..]
func splitToNeedles(s string) []string {
	needles := []string{s}
	for i := 0; i < len(s); i++ {
		needles = append(needles, s[:i])
		needles = append(needles, s[i:])
		needles = append(needles, s[i:i+1])

		// value by N
		for n := 2; n < int(math.Floor(float64(len(s))/2.0)); n++ {
			if i+n < len(s) {
				needles = append(needles, s[i:i+n])
			}
		}

	}
	return needles
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

func check(plain, hash, checkHash string) {
	log.Printf("> %32s\t %s", checkHash, printable(plain))
	if strings.Contains(hash, checkHash) {
		fmt.Println("=========================================================================")
		log.Printf("FOUND: [%s]", printable(plain))
		log.Printf("for %s", hash)
		fmt.Println("=========================================================================")
		os.Exit(0)
	}
}

func addSuffixes(arr, suffixes []string) []string {
	for i, s := range arr {
		arr[i] = s

		for _, suf := range suffixes {
			arr = append(arr, s+suf)
		}
	}
	return arr
}

func printable(s string) string {
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}
