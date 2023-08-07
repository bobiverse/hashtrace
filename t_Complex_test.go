package main

import "testing"

func TestComplex(t *testing.T) {
	testcases := []*TestCase{
		{Comment: "Reversed", ExpectedHash: "5330b35d7c45373", Data: []string{"fire", "truck"}, Separators: []string{"|"}}, // firetruck|kcurterif
		{Comment: "Double hashed", ExpectedHash: "d8266828b55b6c4", Data: []string{"fire", "truck"}},                       // double hash md5(md5(...))
		{Comment: "Tripple hashed", ExpectedHash: "d4febc7be44008e5", Data: []string{"fire", "truck"}},                     // tripple hash md5(md5(md5(...)))
		{Comment: "MD5 to SHA256", ExpectedHash: "c9e5982a4b3c97c4bd178fc4d423a", Data: []string{"fire", "truck"}},         // hash sha256((...))
	}

	for _, tc := range testcases {
		t.Run(tc.Comment, func(t *testing.T) {

			_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
			if err != nil {
				t.Fatalf("%s: `%s` should be cracked by given plain `%q` data", tc.Comment, tc.ExpectedHash, tc.Data)
			}

		})
	}

}
