package bits

/**
 * Set alphabet of words
 */

import "strings"

//var allowedCharacters = "abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -"
var allowedCharacters = "abcdefghijklmnopqrstuvwxyz "
var mapCharToUint = getCharToUintMap(allowedCharacters)
var mapUintToChar = getUintToCharMap(mapCharToUint)

/**
 * Write the data for each node, call getDataBits() to calculate how many bits
 * for one node.
 * 1 bit stores the "final" indicator. The other bits store one of the
 * characters of the alphabet.
 */
var dataBits = getDataBits(allowedCharacters)

func SetAllowedCharacters(alphabet string) {
	allowedCharacters = alphabet
	mapCharToUint = getCharToUintMap(alphabet)
	mapUintToChar = getUintToCharMap(mapCharToUint)
	dataBits = getDataBits(alphabet)
}

func getCharToUintMap(alphabet string) map[string]uint {
	result := map[string]uint{}

	var i uint = 0
	chars := strings.Split(alphabet, "")
	for _, char := range chars {
		result[char] = i
		i++
	}

	return result
}

func getUintToCharMap(c2ui map[string]uint) map[uint]string {
	result := map[uint]string{}
	for k, v := range c2ui {
		result[v] = k
	}
	return result
}

func getDataBits(alphabet string) uint {
	numOfChars := len(strings.Split(alphabet, ""))
	var i uint = 0

	for (1 << i) < numOfChars {
		i++
	}

	// one more bit for the "final" indicator
	return (i + 1)
}
