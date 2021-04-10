package ht

import (
	"fmt"
	"testing"
)

func TestInitialization(t *testing.T) {
	hashtable := New()
	capacity := hashtable.capacity
	length := hashtable.length
	if capacity != initialSize {
		t.Errorf("unexpected output. Received %v", capacity )
	}
	if length != initialSize {
		t.Errorf("unexpected output. Received %v", length )
	}
}

func TestHash(t *testing.T) {
	var tests = []struct{
		key string
		hash uint64
	}{
		{"hello", 0xa430d84680aabd0b},
		{"world",0x4f59ff5e730c8af3 },
		{"example", 0x430b1483c8d66041},
		{"Example", 0x6de3d6e9660857a1},
	}

	hashtable := New()

	for _, value := range tests {
		testName := fmt.Sprintf("hashing %s", value.key)
		t.Run(testName, func(t *testing.T) {
			res := hashtable.hash(value.key)
			if res != value.hash {
				t.Errorf("Expected %#x, received %#x", value.hash, res)
			}
		})
	}
}

//func TestGetter(t *testing.T) {
//	hashtable := New()
//	hashtable.content[4]
//}