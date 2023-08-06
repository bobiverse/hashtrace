package main

import (
	"errors"
	"log"
	"strings"
)

func CrackHash(expectedHash string, dataItems, separators []string) (string, string, error) {
	// 1) Prepare plain and mutated data strings
	needles := PrepareDataNeedles(dataItems, separators)
	//log.Printf("NEEDLES: %d", len(needles))
	log.Printf("%6d needles for %q", len(needles), dataItems)

	N := len(expectedHash)

	// 2) Hash and check
	for _, s := range needles {
		//fmt.Printf(".")

		hashes := []string{
			//hashMD4(s),
			hashMD5(s),
			hashSHA1(s),
			hashSHA256(s),
			hashSHA512(s),
		}

		for _, checkHash := range hashes {

			if compareHash(s, expectedHash, checkHash) {
				return checkHash, s, nil
			}

			if N > len(checkHash) {
				continue
			}

			n := len(checkHash)
			for i := 0; i <= n-N; i++ {
				if i > n {
					continue
				}

				partHash := checkHash[i : i+N]
				if compareHash(s, expectedHash, partHash) {
					return checkHash, s, nil
				}
			}

		}

	}

	return "", "", errors.New("NOT FOUND")
}

func PrepareDataNeedles(dataItems, separators []string) []string {
	var needles []string

	if len(separators) == 0 {
		separators = []string{""}
	}

	// permutations: same length
	for _, perms := range permutations(dataItems) {
		for _, sep := range separators {
			needles = append(needles, strings.Join(perms, sep))
		}
	}

	// combinations: different length
	for _, comb := range combinations(dataItems) {
		for _, perms := range permutations(comb) {
			for _, sep := range separators {
				needles = append(needles, strings.Join(perms, sep))
			}
		}
	}

	// basic transformations
	for _, s := range dataItems {
		needles = append(needles, s)
		needles = append(needles, reverse(s))
		needles = append(needles, splitToNeedles(s)...)
		needles = append(needles, splitToNeedles(reverse(s))...)
	}

	// Lowercase
	for _, s := range needles {
		needles = append(needles, strings.ToLower(s))
	}

	// Uppercase
	for _, s := range needles {
		needles = append(needles, strings.ToUpper(s))
	}

	// LAST) add suffixes
	for _, s := range needles {
		needles = append(needles, s+"\n")
		needles = append(needles, s+"\r\n")
	}

	return uniqueSlice(needles)
}
