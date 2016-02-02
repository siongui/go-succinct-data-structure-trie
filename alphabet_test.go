package bits

import "testing"

func TestAll(t *testing.T) {
	t.Log(mapCharToUint)
	t.Log(mapUintToChar)
	t.Log(dataBits)
	SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	t.Log(mapCharToUint)
	t.Log(mapUintToChar)
	t.Log(dataBits)
	SetAllowedCharacters("abcdefghijklmnopqrstuvwxyz")
}
