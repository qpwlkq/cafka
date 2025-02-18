package common

import (
	"testing"
)

func TestToByte(t *testing.T) {
	cases := []struct {
		name string
		input UNSIGNED_VARINT
		excepted []byte
	} {
		{
			name: "case 1",
			input: 0b1010001010101010,
			excepted: []byte{0b10000010, 0b11000101, 0b00101010},
		},
		{
			name: "case 2",
			input: 0b10010110,
			excepted: []byte{0b10000001, 0b00010110},
		},
	}
	for _, cs := range cases {
		actual := cs.input.ToByte()
		if len(actual) != len(cs.excepted) {
			t.Error("length not equal")
		}
		for i, a := range actual {
			if a != cs.excepted[i] {
				t.Errorf("test failed! excepted: %v, actual: %v", cs.excepted, actual)
			}
		}
	}
}