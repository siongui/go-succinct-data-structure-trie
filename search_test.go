package bits

import "testing"

func TestSearch(t *testing.T) {
	te := Trie{}
	te.Init()
	insertNotInAlphabeticalOrder(&te)
	teData := te.Encode()
	rd := CreateRankDirectory(teData, te.GetNodeCount()*2+1, L1, L2)

	ft := FrozenTrie{}
	ft.Init(teData, rd.GetData(), te.GetNodeCount())

	t.Log(ft.GetSuggestedWords("a", 10))
	t.Log(ft.GetSuggestedWords("b", 10))
	t.Log(ft.GetSuggestedWords("h", 10))
}
