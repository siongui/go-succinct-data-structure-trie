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
package main

import "fmt"

// Configure the bit writing and reading functions to work natively in BASE-64
// encoding. That way, we don't have to convert back and forth to bytes.

var BASE64 =
"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"


/**
    The width of each unit of the encoding, in bits. Here we use 6, for base-64
    encoding.
 */
var W = 6

/**
    Returns the character unit that represents the given value. If this were
    binary data, we would simply return id.
 */
func CHR(id int) string {
	return BASE64[id:id+1]
}


/**
    Returns the decimal value of the given character unit.
 */
var BASE64_CACHE = map[string]int{
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


func ORD(ch string) int {
	// Used to be: return BASE64.indexOf(ch);
	return BASE64_CACHE[ch]
}


/**
    Fixed values for the L1 and L2 table sizes in the Rank Directory
*/
var L1 = 32*32
var L2 = 32


func main() {
	fmt.Println("W:", W)
	fmt.Println("len(BASE64):", len(BASE64))
	fmt.Println("len(BASE64_CACHE):", len(BASE64_CACHE))
	fmt.Println("CHR(0):", CHR(0))
	fmt.Println("ORD(\"A\"):", ORD("A"))
	fmt.Println("CHR(1):", CHR(1))
	fmt.Println("ORD(\"B\"):", ORD("B"))
	fmt.Println("CHR(63):", CHR(63))
	fmt.Println("ORD(\"_\"):", ORD("_"))
	fmt.Println("L1:", L1)
	fmt.Println("L2:", L2)
}
