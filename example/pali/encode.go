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

func saveTrie(t bits.Trie) (err error) {
	rd := bits.CreateRankDirectory(t.Encode(), t.GetNodeCount()*2+1, bits.L1, bits.L2)
	td := TrieData{
		EncodedData:       t.Encode(),
		NodeCount:         t.GetNodeCount(),
		RankDirectoryData: rd.GetData(),
	}

	b, err := json.Marshal(td)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("trie.json", b, 0644)
	return
}

func insertNotInAlphabeticalOrder(te *bits.Trie) {
	te.Insert("sacca")
	te.Insert("ariya")
	te.Insert("saccavācā")
	te.Insert("dhammaṃ")
	te.Insert("buddho")
	te.Insert("viharati")
}

func main() {
	// Set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// Note that you must include space " " in your alphabet if you do not
	// use the default alphabet.
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

	err := saveTrie(te)
	if err != nil {
		panic(err)
	}
}
