package main

import "testing"

func TestTrickyHash(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "b5324104beefb851ff", Data: []string{"a", "b", "c", "d", "e", "f"}, Separators: nil}, // hashed "ac" but given data more
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
