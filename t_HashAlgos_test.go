package main

import "testing"

func TestHashAlgos(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "66a525bf4244d4d403e2f0a86131d243638721c2", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},       // sha1
		{ExpectedHash: "e697ecb2cf692331aab084404d", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                     // sha256
		{ExpectedHash: "ac83a86ede6958600a168b1b60c3dbb6552d873d3b0612", Data: []string{"my", "fire", "truck", "!"}, Separators: nil}, // sha512
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
