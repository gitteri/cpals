package convert

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// HexToB64 converts a hex string to a base64 string
func HexToB64(s string) (string, error){
	var b64 string
	hexBytes, err := hex.DecodeString(s)
	if err != nil {
		return b64, fmt.Errorf("decoding hex failed: %w", err)
	}

	b64 = base64.StdEncoding.EncodeToString(hexBytes)
	return b64, nil
}
