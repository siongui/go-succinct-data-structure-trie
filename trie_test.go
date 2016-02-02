package bits

import "testing"

func insertInAlphabeticalOrder(te *Trie) {
	te.Insert("alphapha")
	te.Insert("apple")
	te.Insert("hello")
	te.Insert("jello")
	te.Insert("lamp")
	te.Insert("orange")
	te.Insert("quiz")
}

func insertNotInAlphabeticalOrder(te *Trie) {
	te.Insert("apple")
	te.Insert("orange")
	te.Insert("alphapha")
	te.Insert("lamp")
	te.Insert("hello")
	te.Insert("jello")
	te.Insert("quiz")
}

func TestTrie(t *testing.T) {
	te := Trie{}
	te.Init()
	insertInAlphabeticalOrder(&te)
	teData := te.Encode()
	t.Log(teData)
	t.Log(te.GetNodeCount())
	if teData != "v2qqqqqqqpIUjQA5JZyBZ4ggCKh55ZZgBA5ZZd5vIEl1wx8g8A" {
		t.Error("Expected v2qqqqqqqpIUjQA5JZyBZ4ggCKh55ZZgBA5ZZd5vIEl1wx8g8A, got ", teData)
	}
	if te.GetNodeCount() != 37 {
		t.Error("Expected 37, got ", te.GetNodeCount())
	}
	rd := CreateRankDirectory(teData, te.GetNodeCount()*2+1, L1, L2)
	if rd.GetData() != "BMIg" {
		t.Error("Expected BMIg, got ", rd.GetData())
	}
	t.Log(rd.GetData())
}
