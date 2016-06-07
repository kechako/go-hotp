package hotp

import "testing"

var key = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30}

func TestGen(t *testing.T) {
	hotp := NewHOTP(key)
	if hotp == nil {
		t.Error("NewHOTP returned nil")
	}

	var expect int
	var code int
	var err error

	// counter: 0, digit: 6
	expect = 755224
	code, err = hotp.Gen(0, 6)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}

	// counter: 0, digit: 8
	expect = 84755224
	code, err = hotp.Gen(0, 8)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}

	// counter: 4, digit: 6
	expect = 338314
	code, err = hotp.Gen(4, 6)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}

	// counter: 4, digit: 8
	expect = 40338314
	code, err = hotp.Gen(4, 8)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}

	// counter: 9, digit: 6
	expect = 520489
	code, err = hotp.Gen(9, 6)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}

	// counter: 9, digit: 8
	expect = 45520489
	code, err = hotp.Gen(9, 8)
	if err != nil {
		t.Error(err)
	}
	if code != expect {
		t.Errorf("want %d, got %d", expect, code)
	}
}
