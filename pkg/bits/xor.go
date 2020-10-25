package bits

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// XorHexStrs will xor two equal length hex strings
func XorHexStrs(s1, s2 string) (string, error) {
	x, err := hex.DecodeString(s1)
	if err != nil {
		return "", fmt.Errorf("decoding hex failed: %w", err)
	}
	y, err := hex.DecodeString(s2)
	if err != nil {
		return "", fmt.Errorf("decoding hex failed: %w", err)
	}

	xorBytes, err := XorEqualLen(x, y)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(xorBytes), nil
}

// XorEqualLen will xor two equal length byte arrays
func XorEqualLen(x, y []byte) ([]byte, error) {
	if len(x) != len(y) {
		return nil, errors.New("can only xor two equal length byte arrays")
	}

	dst := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		dst[i] = x[i] ^ y[i]
	}
	return dst, nil
}
