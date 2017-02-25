// Package hotp implements the HOTP One-Time password algorithm.
package hotp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"math"
)

func truncate(hash []byte) int32 {
	offset := hash[19] & 0x0f

	return (int32(hash[offset])&0x7f)<<24 |
		(int32(hash[offset+1])&0xff)<<16 |
		(int32(hash[offset+2])&0xff)<<8 |
		(int32(hash[offset+3]) & 0xff)
}

// HOTP computes the HOTP One-Time password.
type HOTP struct {
	// Shared secret between client and server.
	Key []byte
}

// New returns a new HOTP computing the HOTP One-Time password.
func New(key []byte) *HOTP {
	return &HOTP{Key: key}
}

// Gen generates a One-Time password based on the counter.
// It trancates the password in digit. If digit <= 0, it doesn't truncate the password.
func (h *HOTP) Gen(counter uint64, digit int) (int32, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 64))
	err := binary.Write(buf, binary.BigEndian, counter)
	if err != nil {
		return 0, err
	}

	mac := hmac.New(sha1.New, h.Key)
	mac.Write(buf.Bytes())
	hash := mac.Sum(nil)

	code := truncate(hash)
	if digit <= 0 {
		return code, nil
	}

	return code % int32(math.Pow10(digit)), nil
}
