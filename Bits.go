/*
 A Succinct Trie for Go

 By Siong-Ui Te
 Released to the public domain.
 translated From:

 A Succinct Trie for Javascript

 By Steve Hanov
 Released to the public domain.

 This file contains functions for creating a succinctly encoded trie structure
 from a list of words. The trie is encoded to a succinct bit string using the
 method of Jacobson (1989). The bitstring is then encoded using BASE-64.

 The resulting trie does not have to be decoded to be used. This file also
 contains functions for looking up a word in the BASE-64 encoded data, in
 O(mlogn) time, where m is the number of letters in the target word, and n is
 the number of nodes in the trie.

 Objects for encoding:

 TrieNode
 Trie
 BitWriter

 Objects for decoding:
 BitString
 FrozenTrieNode
 FrozenTrie

 QUICK USAGE:

 Suppose we let data be some output of the demo encoder:

 var data = {
    "nodeCount": 37,
    "directory": "BMIg",
    "trie": "v2qqqqqqqpIUn4A5JZyBZ4ggCKh55ZZgBA5ZZd5vIEl1wx8g8A"
 };

 var frozenTrie = new FrozenTrie( Data.trie, Data.directory, Data.nodeCount);

 alert( frozenTrie.lookup( "hello" ) ); // outputs true
 alert( frozenTrie.lookup( "kwijibo" ) ); // outputs false

*/
package Bits

import "strings"


// Configure the bit writing and reading functions to work natively in BASE-64
// encoding. That way, we don't have to convert back and forth to bytes.

var BASE64 =
"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"


/**
    The width of each unit of the encoding, in bits. Here we use 6, for base-64
    encoding.
 */
var W uint = 6

/**
    Returns the character unit that represents the given value. If this were
    binary data, we would simply return id.
 */
func CHR(id uint) string {
	return BASE64[id:id+1]
}


/**
    Returns the decimal value of the given character unit.
 */
var BASE64_CACHE = map[string]uint{
	"A" : 0, "B" : 1, "C" : 2, "D" : 3, "E" : 4, "F" : 5, "G" : 6, "H" : 7,
	"I" : 8, "J" : 9, "K" : 10, "L" : 11, "M" : 12, "N" : 13, "O" : 14,
	"P" : 15, "Q" : 16, "R" : 17, "S" : 18, "T" : 19, "U" : 20, "V" : 21,
	"W" : 22, "X" : 23, "Y" : 24, "Z" : 25, "a" : 26, "b" : 27, "c" : 28,
	"d" : 29, "e" : 30, "f" : 31, "g" : 32, "h" : 33, "i" : 34, "j" : 35,
	"k" : 36, "l" : 37, "m" : 38, "n" : 39, "o" : 40, "p" : 41, "q" : 42,
	"r" : 43, "s" : 44, "t" : 45, "u" : 46, "v" : 47, "w" : 48, "x" : 49,
	"y" : 50, "z" : 51, "0" : 52, "1" : 53, "2" : 54, "3" : 55, "4" : 56,
	"5" : 57, "6" : 58, "7" : 59, "8" : 60, "9" : 61, "-" : 62, "_" : 63,
}


func ORD(ch string) uint {
	// Used to be: return BASE64.indexOf(ch);
	return BASE64_CACHE[ch]
}


/**
    Fixed values for the L1 and L2 table sizes in the Rank Directory
*/
var L1 uint = 32*32
var L2 uint = 32


/**
    The BitWriter will create a stream of bytes, letting you write a certain
    number of bits at a time. This is part of the encoder, so it is not
    optimized for memory or speed.
*/
type BitWriter struct {
	bits	[]uint
}


/**
    Write some data to the bit string. The number of bits must be 32 or
    fewer.
*/
func (bw *BitWriter) Write(data, numBits uint) {
	//for i := (numBits-1); i >= 0; i-- {
	//FIXME: the above commented line will cause infinite loop, why???
	for i := numBits; i > 0; i-- {
		j := i-1
		if (data & ( 1 << j )) != 0 {
			bw.bits = append(bw.bits, 1)
		} else {
			bw.bits = append(bw.bits, 0)
		}
	}
}

/**
    Get the bitstring represented as a javascript string of bytes
*/
func (bw *BitWriter) GetData() string {
	var chars []string
	var b, i uint = 0, 0

	for j := 0; j < len(bw.bits); j++ {
		b = ( b << 1 ) | bw.bits[j]
		i += 1
		if i == W {
			chars = append(chars, CHR(b))
			i = 0
			b = 0
		}
	}

	if i != 0 {
		chars = append(chars, CHR(b << ( W - i )) )
	}

	return strings.Join(chars, "")
}

/**
    Returns the bits as a human readable binary string for debugging
 */
func (bw *BitWriter) GetDebugString(group uint) string {
	var chars []string
	var i uint = 0

	for j := 0; j < len(bw.bits); j++ {
		if bw.bits[j] == 1 {
			chars = append(chars, "1")
		} else {
			chars = append(chars, "0")
		}
		i++
		if i == group {
			chars = append(chars, " ")
			i = 0
		}
	}

	return strings.Join(chars, "")
}
