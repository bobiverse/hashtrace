package main

import "testing"

func TestHashAlgos(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "66a525bf4244d4d403e2f0a86131d243638721c2", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                 // sha1
		{ExpectedHash: "e697ecb2cf692331aab084404d", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                               // sha256
		{ExpectedHash: "ac83a86ede6958600a168b1b60c3dbb6552d873d3b0612", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},           // sha512
		{ExpectedHash: "beab1e82cfb7af00726700705c93778a4ecfdd7f4ea4df9d9478790c", Data: []string{"my", "fire", "truck", "!"}, Separators: nil}, // sha224
		{ExpectedHash: "7e90036c28e5e2fddab142c27f2f9993", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                         // sha384
		{ExpectedHash: "94ed5fefc83ef77ea83229393e035834f85e", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                     // sha3-256
		{ExpectedHash: "a011dd0f873df3432901eeb9a57b8fbaecd5f30377", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},               // sha3-512
		{ExpectedHash: "9475a151859f7f7dac23de725edb926db99b", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                     // sha512-224
		{ExpectedHash: "dba532ea996a756d633aa01cc4e936374ca", Data: []string{"my", "fire", "truck", "!"}, Separators: nil},                      // sha512-256
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
