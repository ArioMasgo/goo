package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Authorize(sign, secret, body, path, method string) bool {
	if method == "POST" || method == "PUT" {
		fmt.Print(Signer(body, secret))
		return sign == Signer(body, secret)
	}

	return sign == Signer(path, secret)
}

func Signer(data, secret string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
