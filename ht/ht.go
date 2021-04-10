package ht

import (
	"errors"
)

const (
	initialSize = 16
	FNVOffset = 14695981039346656037
	FNVPrime = 1099511628211
)


type data struct {
	key string
	value int
}

type hasht struct {
	content []data
	length int
	capacity int
}

func New() *hasht {
	content := make([]data, initialSize, initialSize)
	return &hasht{
		content: content,
		length: len(content),
		capacity: cap(content),
	}

}

func (h *hasht) Get(input string) (int, error) {
	key := h.hash(input)
	index := key & uint64(h.capacity - 1)
	// Linear probing: Loop until a blank is found.
	for ; h.content[index].key != ""; {
		if h.content[index].key == input {
			return h.content[index].value, nil
		}

		// loop around if we cross capacity.
		index++
		if int(index) >= h.capacity {
			index = 0
		}
	}

	// If a blank was found, then the key is not in the map.
	return -1, errors.New("input not found")
}

func (h *hasht) Set(input data) error {
	if input.key == "" {
		return errors.New("don't do this to me")
	}

	if h.length >= h.capacity / 2 {
		h.expand()
	}

	key := h.hash(input.key)
	index := key & uint64(h.capacity - 1)

	for ; h.content[index].key != ""; {
		// if we found the key, update the value.
		if h.content[index].key == input.key {
			h.content[index].value = input.value
			return nil
		}

		index++
		if int(index) >= h.capacity {
			index = 0
		}
	}

	// If we find an empty slot, add the key
	h.content[index] = input
	h.length++

	return nil
}

func (h *hasht) Length() int {
	return h.length
}

// hash generates a 64-bit FNV-1A key.
func (h *hasht) hash(input string) uint64 {
	var hash uint64 = FNVOffset
	for _, char := range input {
		hash ^= uint64(char)
		hash *= FNVPrime
	}
	return hash
}

// expand will increase the size of the hashset.
func (h *hasht) expand() error {
	return nil
}