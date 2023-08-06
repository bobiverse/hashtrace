package main

import (
	"errors"
)

func CrackHash(expectedHash string, dataItems, separators []string) (string, string, error) {
	// 1) Prepare plain and mutated data strings
	needles := PrepareDataNeedles(dataItems)

	N := len(expectedHash)

	// 2) Hash and check
	for _, s := range needles {

		hashes := []string{
			//hashMD4(s),
			hashMD5(s),
			//hashSHA1(s),
			//hashSHA256(s),
			//hashSHA512(s),
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

func PrepareDataNeedles(dataItems []string) []string {
	var needles []string

	for _, s := range dataItems {
		needles = append(needles, s)
		needles = append(needles, reverse(s))
		needles = append(needles, splitToNeedles(s)...)
		needles = append(needles, splitToNeedles(reverse(s))...)
	}

	//// glue
	//for _, s := range data
	//

	// END) add suffixes
	for _, s := range needles {
		needles = append(needles, s+"\n")
		needles = append(needles, s+"\r\n")
	}

	return uniqueSlice(needles)
}
