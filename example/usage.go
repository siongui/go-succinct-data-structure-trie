package main

import (
	bits "github.com/siongui/go-succinct-data-structure-trie"
)

func insertNotInAlphabeticalOrder(te *bits.Trie) {
	te.Insert("apple")
	te.Insert("orange")
	te.Insert("alphapha")
	te.Insert("lamp")
	te.Insert("hello")
	te.Insert("jello")
	te.Insert("quiz")
}

func main() {
	// optional: set alphabet of words
	//bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// Note that you must include space " " in your alphabet if you do not use the
	// default alphabet.
	// default alphabet is [a-z ], i.e.,
	// bits.SetAllowedCharacters("abcdefghijklmnopqrstuvwxyz ")

	// encode: build succinct trie
	te := bits.Trie{}
	te.Init()
	// encode: insert words
	insertNotInAlphabeticalOrder(&te)
	// encode: trie encoding
	teData := te.Encode()
	println(teData)
	println(te.GetNodeCount())
	// encode: build cache for quick lookup
	rd := bits.CreateRankDirectory(teData, te.GetNodeCount()*2+1, bits.L1, bits.L2)
	println(rd.GetData())

	// decode: build frozen succinct trie
	ft := bits.FrozenTrie{}
	ft.Init(teData, rd.GetData(), te.GetNodeCount())

	// decode: look up words
	println(ft.Lookup("apple"))
	println(ft.Lookup("appl"))
	println(ft.Lookup("applee"))

	// decode: words suggestion (find words that start with "prefix")
	// find words starts with "a", max number of returned words is 10
	for _, word := range ft.GetSuggestedWords("a", 10) {
		println(word)
	}
}
