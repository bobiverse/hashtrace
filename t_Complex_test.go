package main

import "testing"

func TestComplex(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "5330b35d7c45373", Data: []string{"fire", "truck"}},  // firetruck|kcurterif
		{ExpectedHash: "d8266828b55b6c4", Data: []string{"fire", "truck"}},  // double hash md5(md5(...))
		{ExpectedHash: "d8266828b55b6c45", Data: []string{"fire", "truck"}}, // tripple hash md5(md5(md5(...)))
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
