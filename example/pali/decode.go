package main

import (
	"encoding/json"
	"io/ioutil"

	bits "github.com/siongui/go-succinct-data-structure-trie"
)

type TrieData struct {
	EncodedData       string
	NodeCount         uint
	RankDirectoryData string
}

func loadTrie(filePath string) (td TrieData, err error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &td)
	return
}

func main() {
	// Set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// Note that you must include space " " in your alphabet if you do not
	// use the default alphabet.
	// default alphabet is [a-z ], i.e.,
	// bits.SetAllowedCharacters("abcdefghijklmnopqrstuvwxyz ")

	td, err := loadTrie("trie.json")
	if err != nil {
		panic(err)
	}

	println(td.EncodedData)
	println(td.NodeCount)
	println(td.RankDirectoryData)

	// decode: build frozen succinct trie
	ft := bits.FrozenTrie{}
	ft.Init(td.EncodedData, td.RankDirectoryData, td.NodeCount)

	// decode: look up words
	println(ft.Lookup("sacca"))
	println(ft.Lookup("sacc"))
	println(ft.Lookup("dhamma"))

	// decode: words suggestion (find words that start with "prefix")
	// find words starts with "a", max number of returned words is 10
	for _, word := range ft.GetSuggestedWords("a", 10) {
		println(word)
	}
}
