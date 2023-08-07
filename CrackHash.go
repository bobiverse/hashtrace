package main

import (
	"errors"
	"log"
	"net/url"
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

		hashes := makeHashes(s)

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

	// glue all data items in one string and generate needles
	if len(dataItems) > 1 {
		needles = append(needles, PrepareDataNeedles([]string{strings.Join(dataItems, "")}, separators)...)
	}

	// add reversed too
	for _, s := range dataItems {
		dataItems = append(dataItems, reverse(s))
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
		needles = append(needles, splitToNeedles(s)...)
	}

	// URL encoded/decoded
	for _, s := range needles {
		needles = append(needles, url.QueryEscape(s)) // " " => "+"
		needles = append(needles, url.PathEscape(s))  // " " => "%20"

		sDecoded, _ := url.QueryUnescape(s)
		needles = append(needles, sDecoded)

		sDecoded, _ = url.PathUnescape(s)
		needles = append(needles, sDecoded)
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
