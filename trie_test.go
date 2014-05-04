package trie

import (
	"strconv"
	"testing"
)

func Test_ExistsOrAdd_1(t *testing.T) {
	r := NewTrie()
	r.ExistsOrAdd("andre")
	if ! r.ExistsOrAdd("andre") {
		t.Error("ExistsOrAdd_1: Add failed!")
	} else
	{
		t.Log("ExistsOrAdd_1: Succesful")
	}
}

func Test_ExistsOrAdd_2(t *testing.T) {
	r := NewTrie()
	r.ExistsOrAdd("andre")
	if r.ExistsOrAdd("and") {
		t.Error("ExistsOrAdd_2: Add failed!")
	} else
	{
		t.Log("ExistsOrAdd_2: Succesful")
	}
}

func Test_ExistsOrAdd_3(t *testing.T) {
	r := NewTrie()
	for i := 1; i < 1000000; i++ {
		r.ExistsOrAdd(strconv.Itoa(i))
	}
	if ! r.ExistsOrAdd("666666") {
		t.Error("ExistsOrAdd_3: Add failed!")
	} else
	{
		t.Log("ExistsOrAdd_3: Succesful")
	}
}
