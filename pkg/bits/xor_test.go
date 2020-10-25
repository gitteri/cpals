package bits

import (
	"strings"
	"testing"
)

func TestXorHexStrs(t *testing.T) {
	cs := []struct {
		name        string
		hex1        string
		hex2        string
		expectedHex string
		errLike     string
	}{
		{
			name:        "nil hex",
			hex1:        "",
			hex2:        "",
			expectedHex: "",
		},
		{
			name:        "given cryptopals case",
			hex1:        "1c0111001f010100061a024b53535009181c",
			hex2:        "686974207468652062756c6c277320657965",
			expectedHex: "746865206b696420646f6e277420706c6179",
		},
		{
			name:    "invalid hex lengths",
			hex1:    "123456",
			hex2:    "1234",
			errLike: "can only xor two equal length",
		},
		{
			name:    "invalid hex encoding",
			hex1:    "x*s",
			hex2:    "123",
			errLike: "decoding hex failed",
		},
	}

	for _, cx := range cs {
		c := cx
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			hex, err := XorHexStrs(c.hex1, c.hex2)
			if c.errLike != "" {
				if err == nil {
					t.Fatal("expected error but got none")
				}
				if !strings.Contains(err.Error(), c.errLike) {
					t.Fatal("expected error", c.errLike, "but got err", err.Error())
				}
			}

			if hex != c.expectedHex {
				t.Fatal("expected hex", c.expectedHex, "but got hex %s", hex)
			}
		})
	}
}
