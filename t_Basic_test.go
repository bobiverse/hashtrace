package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	testcases := map[string]string{
		"68b329da9893e34099c7d8ad5cb9c940": "",
		"e1c06d85ae7b8b032bef47e42e4c08f9": "\n",
		"60b725f10c9c85c70d97880dfe8191b3": "a",
		"90909a7058bdf9377ed69a157678713c": "Šrubberī", // Š
	}

	for inHash, inData := range testcases {
		_, _, err := CrackHash(inHash, []string{inData}, nil)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", inHash, inData)
		}
	}

}