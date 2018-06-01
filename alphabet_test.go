package bits

import (
	"testing"
)

func TestAlphabet(t *testing.T) {
	if len(mapCharToUint) != 27 {
		t.Error("len(mapCharToUint) != 27")
		t.Log(mapCharToUint)
	}
	if len(mapUintToChar) != 27 {
		t.Error("len(mapUintToChar) != 27")
		t.Log(mapUintToChar)
	}
	if dataBits != 6 {
		t.Error("dataBits != 6")
		t.Log(dataBits)
	}

	SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")

	if len(mapCharToUint) != 39 {
		t.Error("len(mapCharToUint) != 39")
		t.Log(mapCharToUint)
	}
	if len(mapUintToChar) != 39 {
		t.Error("len(mapUintToChar) != 39")
		t.Log(mapUintToChar)
	}
	if dataBits != 7 {
		t.Error("dataBits != 7")
		t.Log(dataBits)
	}

	SetAllowedCharacters("abcdefghijklmnopqrstuvwxyz ")
	if dataBits != 6 {
		t.Error("dataBits != 6")
		t.Log(dataBits)
	}
}
