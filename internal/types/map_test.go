package types

import (
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name, key, value string
	}{
		{
			name:  "Test with array",
			key:   "a" + KeySep + "aa",
			value: "1" + ValSep + "2",
		},
		{
			name:  "Test with map",
			key:   "b" + KeySep + "bb",
			value: "3",
		},
		{
			name:  "Test with string",
			key:   "c",
			value: "4",
		},
		{
			name:  "Test with nil",
			key:   "e",
			value: "",
		},
	}

	m, err := NewMap(`{
		"a": [{
			"aa": "1"
		}, {
			"aa": "2"
		}],
		"b": {
			"bb": "3"
		},
		"c": "4"
	}`)

	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v := m.GetString(tt.key)

			if v != tt.value {
				t.Fatal("Value mismatch")
			}
		})
	}
}
