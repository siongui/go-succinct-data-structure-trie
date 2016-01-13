package Bits

import "testing"

func TestBASE64(t *testing.T) {
	t.Log("W:", W)
	t.Log("len(BASE64):", len(BASE64))
	if len(BASE64) != 64 {
		t.Error("len(BASE64) Expected 64, got ", len(BASE64))
	}
	t.Log("len(BASE64_CACHE):", len(BASE64_CACHE))
	if len(BASE64_CACHE) != 64 {
		t.Error("len(BASE64_CACHE) Expected 64, got ", len(BASE64_CACHE))
	}
	t.Log("CHR(0):", CHR(0))
	if CHR(0) != "A" {
		t.Error("CHR(0) Expected \"A\", got ", CHR(0))
	}
	t.Log("ORD(\"A\"):", ORD("A"))
	if ORD("A") != 0 {
		t.Error("ORD(\"A\") Expected 0, got ", ORD("A"))
	}
	t.Log("CHR(1):", CHR(1))
	t.Log("ORD(\"B\"):", ORD("B"))
	t.Log("CHR(63):", CHR(63))
	if CHR(63) != "_" {
		t.Error("CHR(63) Expected \"_\", got ", CHR(63))
	}
	t.Log("ORD(\"_\"):", ORD("_"))
	if ORD("_") != 63 {
		t.Error("ORD(\"_\") Expected 63, got ", ORD("_"))
	}
	t.Log("L1:", L1)
	t.Log("L2:", L2)
}

func TestBitWriter(t *testing.T) {
	bw := BitWriter{}
	bw.Write(3, 2)
	if bw.GetDebugString(3) != "11" {
		t.Error("Expected 11, got ", bw.GetDebugString(3))
	}
	if bw.GetData() != "w" {
		t.Error("Expected w, got ", bw.GetData())
	}
	bw.Write(0, 3)
	if bw.GetData() != "w" {
		t.Error("Expected w, got ", bw.GetData())
	}
	bw.Write(2, 2)
	if bw.GetData() != "xA" {
		t.Error("Expected xA, got ", bw.GetData())
	}
	t.Log(bw)
	t.Log(bw.GetData())
	t.Log(bw.GetDebugString(3))
	if bw.GetDebugString(3) != "110 001 0" {
		t.Error("Expected 110 001 0, got ", bw.GetDebugString(3))
	}
}

func TestBitString(t *testing.T) {
	bs := BitString{}
	bs.Init("88kj5w_6phb")
	t.Log(bs)
	if bs.Rank(5) != 4 {
		t.Error("Expected 4, got ", bs.Rank(5))
	}
	if bs.Rank(24) != 14 {
		t.Error("Expected 14, got ", bs.Rank(24))
	}
	if bs.Rank(37) != 21 {
		t.Error("Expected 21, got ", bs.Rank(37))
	}
	if bs.Rank(55) != 33 {
		t.Error("Expected 33, got ", bs.Rank(55))
	}
	if bs.Rank(65) != 38 {
		t.Error("Expected 38, got ", bs.Rank(65))
	}
	// FIXME??: bs.Rank(66) fails the test
	if bs.Get(5,7) != 60 {
		t.Error("Expected 60, got ", bs.Get(5,7))
	}
	if bs.Get(7,13) != 7314 {
		t.Error("Expected 7314, got ", bs.Get(7,13))
	}
	if bs.Get(0,5) != 30 {
		t.Error("Expected 30, got ", bs.Get(0,5))
	}
	if bs.Get(3,3) != 4 {
		t.Error("Expected 4, got ", bs.Get(3,3))
	}
	if bs.Get(33,17) != 16362 {
		t.Error("Expected 16362, got ", bs.Get(33,17))
	}
	if bs.Count(0,17) != 10 {
		t.Error("Expected 10, got ", bs.Count(0,17))
	}
	if bs.Count(7,2) != 2 {
		t.Error("Expected 2, got ", bs.Count(7,2))
	}
	if bs.Count(56,9) != 4 {
		t.Error("Expected 4, got ", bs.Count(56,9))
	}
	if bs.Count(12,1) != 1 {
		t.Error("Expected 1, got ", bs.Count(12,1))
	}
	if bs.Count(5,7) != 4 {
		t.Error("Expected 4, got ", bs.Count(5,7))
	}
}

func TestRankDirectory(t *testing.T) {
	rd := CreateRankDirectory("1wnc2bxhbx7mkbgnpwq7vtlub7p6pkls42lvie9j1ekcpt0zytrdl67enescolwex7aumq4imywstrpktbvxy0rp61nnonj9grdf", 400, L1, L2)
	t.Log(rd)
	if rd.directory.GetData() != "BIJA0EcXBsH4kykLgzjc" {
		t.Error("Expected BIJA0EcXBsH4kykLgzjc, got ", rd.directory.GetData())
	}
	if rd.directory.length != 120 {
		t.Error("Expected 120, got ", rd.directory.length)
	}
	if rd.Rank(1, 200) != 113 {
		t.Error("Expected 113, got ", rd.Rank(1, 200))
	}
	if rd.Rank(0, 100) != 47 {
		t.Error("Expected 47, got ", rd.Rank(0, 100))
	}
	if rd.Select(1, 134) != 233 {
		t.Error("Expected 233, got ", rd.Rank(1, 134))
	}
	if rd.Select(0, 77) != 178 {
		t.Error("Expected 178, got ", rd.Rank(0, 77))
	}
}

func TestTrie(t *testing.T) {
	te := Trie{}
	te.Init()
	te.Insert("apple")
	te.Insert("orange")
	te.Insert("alphapha")
	te.Insert("lamp")
	te.Insert("hello")
	te.Insert("jello")
	te.Insert("quiz")
	teData := te.Encode()
	t.Log(teData)
	t.Log(te.GetNodeCount())
	if teData != "v6qqqqqqqpRQp-AcAWOSgeiWAIIoeAeYWWQWaPeWXzIMBddIePA" {
		t.Error("Expected v6qqqqqqqpRQp-AcAWOSgeiWAIIoeAeYWWQWaPeWXzIMBddIePA, got ", teData)
	}
	if te.GetNodeCount() != 38 {
		t.Error("Expected 38, got ", te.GetNodeCount())
	}
	rd := CreateRankDirectory(teData, te.GetNodeCount() * 2 + 1, L1, L2)
	if rd.GetData() != "BMIg" {
		t.Error("Expected BMIg, got ", rd.GetData())
	}
	t.Log(rd.GetData())
}
