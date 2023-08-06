package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	testcases := map[string]string{
		"68b329da9893e34099c7d8ad5cb9c940":    "",
		"e1c06d85ae7b8b032bef47e42e4c08f9":    "\n",
		"60b725f10c9c85c70d97880dfe8191b3":    "a",
		"90909a7058bdf9377ed69a157678713c":    "Šrubberī", // md5 Š
		"78641538fedf777":                     "ĀĀĀ",      // md5 ā
		"1d7d8700a1c181a7d6c823d059f3c9ffXXX": "čau",      // md5 ČAU
	}

	for inHash, inData := range testcases {
		_, _, err := CrackHash(inHash, []string{inData}, nil)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", inHash, inData)
		}
	}

}

func TestBasicSeparators(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "7e94122", Data: []string{"a", "b", "c", "d"}, Separators: []string{"", "|"}},
		{ExpectedHash: "52a4e09a7483", Data: []string{"water", "truck", "bus", "fire"}, Separators: []string{"", "|"}},
		{ExpectedHash: "c026a88b6", Data: []string{"fake", "truck", "TRUNK", "FIRE"}, Separators: nil},
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
