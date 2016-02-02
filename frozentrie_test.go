package bits

import "testing"

func TestLookup(t *testing.T) {
	te := Trie{}
	te.Init()
	insertNotInAlphabeticalOrder(&te)
	teData := te.Encode()
	rd := CreateRankDirectory(teData, te.GetNodeCount()*2+1, L1, L2)

	ft := FrozenTrie{}
	ft.Init(teData, rd.GetData(), te.GetNodeCount())

	if ft.Lookup("apple") != true {
		t.Error("apple")
	}
	if ft.Lookup("appl") != false {
		t.Error("appl")
	}
	if ft.Lookup("applea") != false {
		t.Error("applea")
	}
	if ft.Lookup("orange") != true {
		t.Error("orange")
	}
	if ft.Lookup("lamp") != true {
		t.Error("lamp")
	}
	if ft.Lookup("hello") != true {
		t.Error("hello")
	}
	if ft.Lookup("jello") != true {
		t.Error("jello")
	}
	if ft.Lookup("quiz") != true {
		t.Error("quiz")
	}
	if ft.Lookup("quize") != false {
		t.Error("quize")
	}
	if ft.Lookup("alphaph") != false {
		t.Error("alphaph")
	}
	if ft.Lookup("alphapha") != true {
		t.Error("alphapha")
	}
}
