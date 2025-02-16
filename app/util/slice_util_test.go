package util

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		Name string
		input, except []byte
	} {
		{
			Name: "case 1",
			input: []byte{254, 1, 2, 3, 4, 255},
			except: []byte{255, 4, 3, 2, 1, 254},
		},
		{
			Name: "case 2",
			input: []byte{},
			except: []byte{},
		},
	}

	for _, c := range cases {
		result := Reverse(c.input)
		if len(c.except) != len(result) {
			t.Error("length not equal")
		}
		for i := 0; i < len(c.except); i++ {
			if c.except[i] != result[i] {
				t.Errorf("test failed except: %v, actual: %v", c.except, result)
				break
			}
		}
	}
}