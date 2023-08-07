package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
	"time"
)

const (
	ClrReset  = "\033[0m"
	ClrBlack  = "\033[30m"
	ClrRed    = "\033[31m"
	ClrGreen  = "\033[32m"
	ClrYellow = "\033[33m"
	ClrBlue   = "\033[34m"
	ClrPurple = "\033[35m"
	ClrCyan   = "\033[36m"
	ClrWhite  = "\033[37m"

	ClrBgBlack  = "\033[40m"
	ClrBgRed    = "\033[41m"
	ClrBgGreen  = "\033[42m"
	ClrBgYellow = "\033[43m"
	ClrBgBlue   = "\033[44m"
	ClrBgPurple = "\033[45m"
	ClrBgCyan   = "\033[46m"
	ClrBgWhite  = "\033[47m"
)

func main() {
	var expectedHash string
	var dataItems flagSlice
	var separators flagSlice
	flag.StringVar(&expectedHash, "hash", "", "target expectedHash to crack")
	flag.Var(&dataItems, "data", "data points possibly used to expectedHash")
	flag.Var(&separators, "sep", "possible separator")
	flag.Parse()

	if expectedHash == "" || dataItems.IsEmpty() {
		log.Fatalf("Usage:\n\t./hashtrace -hash='90b76b4e' -data='firetruck'\n\t./hashtrace -hash='76b4e78167' -data=truck -data=my --data=fire -sep=','\n")
	}

	// Separators
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	detectedSeparators := re.FindAllString(dataItems.String(), -1)
	if separators.IsEmpty() {
		separators = append(separators, detectedSeparators...)
		separators = append(separators, []string{"|", ",", ";"}...)
	}
	separators = append(separators, "")
	separators = uniqueSlice(separators)

	// Used variables
	fmt.Println(strings.Repeat(".", 80))
	fmt.Println(">> " + expectedHash)
	fmt.Println(">> " + dataItems.String() + "")
	if len(separators) > 0 {
		fmt.Println(printable(fmt.Sprintf(">> Separators %q", separators)))
	}
	fmt.Println(strings.Repeat(".", 80))

	// Crack!
	foundHash, foundPlain, err := CrackHash(expectedHash, dataItems, separators)
	if err != nil {
		log.Fatalf("NO LUCK: %v: %s", dataItems, err)
		return
	}

	fmt.Println("=========================================================================")
	log.Printf("FOUND AT %s", time.Now())
	fmt.Printf("PLAIN:\t%s%q%s\n", ClrGreen, foundPlain, ClrReset)
	fmt.Printf("HASH:\t %s\n", ClrGreen+expectedHash+ClrReset)
	if foundHash != expectedHash {
		// fmt.Printf("\t\t    FULL:\t %s\n", foundHash)
		fmt.Printf("FULL:\t %s\n", strings.ReplaceAll(foundHash, expectedHash, ClrBgGreen+ClrBlack+expectedHash+ClrReset))
	}

	fmt.Println("=========================================================================")

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

func compareHash(plain, hash, checkHash string) bool {
	checkHashes := []string{
		checkHash,
		reverse(checkHash),
	}

	for _, ch := range checkHashes {
		if strings.Contains(hash, ch) {
			return true
		}
	}

	return false
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

func mutateWithSeparator(combinations []string, splitBySep, sep string) []string {
	var mutations []string
	for _, s := range combinations {
		time.Sleep(1 * time.Millisecond)
		parts := strings.Split(s, splitBySep)
		mutations = append(mutations, strings.Join(parts, sep))
		mutations = append(mutations, strings.Join(parts, sep)+sep)
		mutations = append(mutations, sep+strings.Join(parts, sep))
		mutations = append(mutations, sep+strings.Join(parts, sep)+sep)
	}
	return mutations
}
