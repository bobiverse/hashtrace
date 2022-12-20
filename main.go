package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	var hash string
	var data string
	var separators flagSlice
	flag.StringVar(&hash, "hash", "", "target hash to crack")
	flag.StringVar(&data, "data", "", "data possibly used to hash")
	flag.Var(&separators, "separators", "possible separator")
	flag.Parse()

	if hash == "" || data == "" {
		log.Fatalf("Usage:\n\t./hashguess -hash='90b76b4e' -data='firetruck'\n")
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	detectedSeparators := re.FindAllString(data, -1)
	if separators.IsEmpty() {
		separators = append(separators, detectedSeparators...)
		separators = append(separators, []string{" ", "|", ",", ";", "_", "-", "/", "\t", "\n", "\r"}...)
	}

	// Used variables
	fmt.Println(strings.Repeat(".", 80))
	fmt.Println(">> " + hash)
	fmt.Println(">> " + data + "")
	if len(separators) > 0 {
		fmt.Println(printable(fmt.Sprintf(">> Separators %v", separators)))
	}
	fmt.Println(strings.Repeat(".", 80))

	// generate needles
	var needles []string
	needles = append(needles, data)
	needles = append(needles, reverse(data))
	needles = append(needles, splitToNeedles(data)...)
	needles = append(needles, splitToNeedles(reverse(data))...)

	// split to parts based on separator
	var sepNeedles []string
	for _, sep := range separators {
		separr := splitBySeparator(data, sep)
		needles = append(needles, separr...)
		for _, detectedSep := range detectedSeparators {
			if detectedSep == sep {
				continue
			}
			sepNeedles = append(sepNeedles, mutateWithSeparator(needles, detectedSep, sep)...)
			fmt.Println(detectedSep, sep)
		}
	}
	needles = append(needles, sepNeedles...)

	// final touch
	needles = append(needles, addSuffixes(needles, []string{"\n", "\r", "\r\n", "\n\r"})...)
	needles = uniqueSlice(needles)

	N := len(hash)
	for _, s := range needles {

		if len(s) > 5 && strings.Contains(s, "|") {
			fmt.Println("########### " + s)
		}

		// hashesh to check
		hashes := []string{
			hashMD5(s),
		}

		// if source hash is shorter
		for _, checkHash := range hashes {
			check(s, hash, checkHash)

			if N > len(checkHash) {
				continue
			}

			n := len(checkHash)
			for i := 0; i <= n-N; i++ {
				if i > n {
					continue
				}

				// 1.
				partHash := checkHash[i : i+N]
				check(s, hash, partHash)
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

func check(plain, hash, checkHash string) {
	checkHashes := []string{
		checkHash,
		reverse(checkHash),
	}

	for _, ch := range checkHashes {
		if strings.Contains(hash, ch) {
			fmt.Println("=========================================================================")
			log.Printf("FOUND:\t [%s]", printable(plain))
			log.Printf("for hash\t %s", hash)
			fmt.Println("=========================================================================")
			os.Exit(0)
		}
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

func splitBySeparator(s, sep string) []string {
	parts := strings.Split(s, sep)
	all := permutations(parts)

	var combinations []string
	for _, parts := range all {
		combinations = append(combinations, strings.Join(parts, sep))
		combinations = append(combinations, strings.Join(parts, ""))
	}
	return combinations
}

func mutateWithSeparator(combinations []string, splitBySep, sep string) []string {
	var mutations []string
	for _, s := range combinations {
		parts := strings.Split(s, splitBySep)
		mutations = append(mutations, strings.Join(parts, sep))
		mutations = append(mutations, strings.Join(parts, sep)+sep)
		mutations = append(mutations, sep+strings.Join(parts, sep))
		mutations = append(mutations, sep+strings.Join(parts, sep)+sep)
	}
	return mutations
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				tmp := arr[i]
				arr[i] = arr[n-1]
				arr[n-1] = tmp
				continue
			}

			tmp := arr[0]
			arr[0] = arr[n-1]
			arr[n-1] = tmp
		}
	}
	helper(arr, len(arr))
	return res
}
