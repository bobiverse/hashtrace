package main

import "testing"

func TestBasicSeparators(t *testing.T) {
	testcases := []*TestCase{
		{ExpectedHash: "7e94122", Data: []string{"a", "b", "c", "d"}, Separators: []string{"", "|"}},
		{ExpectedHash: "52a4e09a7483", Data: []string{"water", "truck", "bus", "fire"}, Separators: []string{"", "|"}},
		{ExpectedHash: "c026a88b6", Data: []string{"fake", "truck", "TRUNK", "FIRE"}, Separators: nil},
		{ExpectedHash: "a3b3a3ff0601", Data: []string{"domain.xx", "john", "fake"}, Separators: []string{"@"}}, // custom separator
	}

	for _, tc := range testcases {
		_, _, err := CrackHash(tc.ExpectedHash, tc.Data, tc.Separators)
		if err != nil {
			t.Fatalf("`%s` should be cracked by given plain `%q` data", tc.ExpectedHash, tc.Data)
		}
	}

}
