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
