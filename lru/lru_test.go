package lru

import (
	"testing"
	"fmt"
	. "github.com/liunian1004/gds/utils"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(3)

	lru.Put(1, "a")
	lru.Put(2, "b")
	lru.Put(3, "c")
	lru.Put(4, "f")
	lru.Put(5, "g")

	fmt.Printf("%+v\n", lru.Keys())

	for _, value := range lru.m {
		fmt.Printf("%+v\n", value)
	}

	// keys()
	if keys, expected := lru.Keys(), []interface{}{3, 4, 5}; !EqualSlice(keys, expected) {
		t.Errorf("Get %+v Expected %+v", keys, expected)
	}
}

func TestLRU_Get(t *testing.T) {

}

