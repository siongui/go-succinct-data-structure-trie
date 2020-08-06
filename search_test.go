package bits

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	te := Trie{}
	te.Init()
	insertNotInAlphabeticalOrder(&te)
	teData := te.Encode()
	rd := CreateRankDirectory(teData, te.GetNodeCount()*2+1, L1, L2)

	ft := FrozenTrie{}
	ft.Init(teData, rd.GetData(), te.GetNodeCount())

	if !reflect.DeepEqual(ft.GetSuggestedWords("a", 10), []string{"apple", "alphapha"}) {
		t.Error(`ft.GetSuggestedWords("a", 10) != []string{"apple", "alphapha"}`)
	}
	if len(ft.GetSuggestedWords("b", 10)) != 0 {
		t.Error(`len(ft.GetSuggestedWords("b", 10)) != 0`)
	}
	if !reflect.DeepEqual(ft.GetSuggestedWords("h", 10), []string{"hello"}) {
		t.Error(`ft.GetSuggestedWords("h", 10) != []string{"hello"}`)
	}
}
