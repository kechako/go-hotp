package hotp

import (
	"reflect"
	"testing"
)

var key = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30}

var hotpTests = []struct {
	inCounter uint64
	inDec     int
	out       int
}{
	{0, 6, 755224},
	{1, 6, 287082},
	{2, 6, 359152},
	{3, 6, 969429},
	{4, 6, 338314},
	{5, 6, 254676},
	{6, 6, 287922},
	{7, 6, 162583},
	{8, 6, 399871},
	{9, 6, 520489},
	{0, 8, 84755224},
	{1, 8, 94287082},
	{2, 8, 37359152},
	{3, 8, 26969429},
	{4, 8, 40338314},
	{5, 8, 68254676},
	{6, 8, 18287922},
	{7, 8, 82162583},
	{8, 8, 73399871},
	{9, 8, 45520489},
}

func TestGen(t *testing.T) {
	hotpExpect := &HOTP{Key: key}
	hotp := New(key)
	if !reflect.DeepEqual(hotp, hotpExpect) {
		t.Errorf("New(%v)\n=> %#v,\nwant %#v", key, hotp, hotpExpect)
	}

	for _, tt := range hotpTests {
		code, err := hotp.Gen(tt.inCounter, tt.inDec)
		if err != nil {
			t.Errorf("Gen(%d, %d) => %v", tt.inCounter, tt.inDec, err)
		} else if code != tt.out {
			t.Errorf("Gen(%d, %d) => %d, want %d", tt.inCounter, tt.inDec, code, tt.out)
		}
	}
}
