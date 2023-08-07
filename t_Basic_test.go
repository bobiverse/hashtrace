package main

import (
	"testing"
)

func TestBasicX(t *testing.T) {
	testcases := map[string]string{
		"68b329da9893e34099c7d8ad5cb9c940":    "",
		"e1c06d85ae7b8b032bef47e42e4c08f9":    "\n",
		"60b725f10c9c85c70d97880dfe8191b3":    "a",
		"90909a7058bdf9377ed69a157678713c":    "Šrubberī",     // md5 one letter `Š`
		"78641538fedf777":                     "ĀĀĀ",          // md5 one letter `ā`
		"1d7d8700a1c181a7d6c823d059f3c9ffXXX": "čau",          // md5 uppercased ČAU
		"50a1154f3e882f7951b":                 "fire%20truck", // md5 "fire truck" // urlencoded
		"b71bfee412137cd0092":                 "fire%20truck", // md5 "fire%20truck" // urlencoded
		"50a1154f3e882f795":                   "fire truck",   // md5 "fire truck" // urlencoded
		"b71bfee412137cd00":                   "fire truck",   // md5 "fire%20truck" // urlencoded
		"cfac1625a2187da60a":                  "firetruck",    // reverse
	}

	for inHash, inData := range testcases {
		_, _, err := CrackHash(inHash, []string{inData}, nil)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", inHash, inData)
		}
	}

}
