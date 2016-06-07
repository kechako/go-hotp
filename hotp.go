package hotp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"math"
)

func truncate(hash []byte) uint32 {
	offset := hash[19] & 0x0f

	return (uint32(hash[offset])&0x7f)<<24 |
		(uint32(hash[offset+1])&0xff)<<16 |
		(uint32(hash[offset+2])&0xff)<<8 |
		(uint32(hash[offset+3]) & 0xff)
}

type HOTP struct {
	Key []byte
}

func NewHOTP(key []byte) *HOTP {
	return &HOTP{Key: key}
}

func (h *HOTP) Gen(counter uint64, digit int) (int, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 64))
	err := binary.Write(buf, binary.BigEndian, counter)
	if err != nil {
		return 0, err
	}

	mac := hmac.New(sha1.New, h.Key)
	mac.Write(buf.Bytes())
	hash := mac.Sum(nil)

	code := truncate(hash)

	return int(code % uint32(math.Pow10(digit))), nil
}
