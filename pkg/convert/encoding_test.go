package convert

import "testing"

func TestHexToB64(t *testing.T) {
	cs := []struct {
		name        string
		hex         string
		expectedB64 string
		expectedErr string
	}{
		{
			name:        "nil hex",
			hex:         "",
			expectedErr: "Could not convert hex to b64",
		},
		{
			name:        "given cryptopals case",
			hex:         "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			expectedB64: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			b64, err := HexToB64(c.hex)
			if c.expectedErr != "" {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if err.Error() != c.expectedErr {
					t.Fatal("expected error", c.expectedErr, "but got err", err.Error())
				}
			}

			if b64 != c.expectedB64 {
				t.Fatal("expected b64", c.expectedB64, "but got b64 %s", b64)
			}
		})
	}
}
