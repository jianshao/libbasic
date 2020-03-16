package datastruct

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	HashTableInvalidOffset = -1
)

type GetOffsetFromKey func(key interface{}, hashMask int) int

type nodeData struct {
	key interface{}
	data interface{}
}

type hashNode struct {
	count int
	nodes []nodeData
	rwLock sync.RWMutex
}

type HashTable struct {
	hashMask int
	getKey GetOffsetFromKey
	hash []hashNode
}

func NewHashTable(hashMask int, getKey GetOffsetFromKey) *HashTable {
	return &HashTable{
		hashMask:hashMask,
		getKey:getKey,
		hash:make([]hashNode, hashMask),
	}
}

func (h *HashTable)Add(key interface{}, data interface{}) (bool, error) {

	if nil == key {
		return false, fmt.Errorf("invalid param")
	}

	offset := h.getKey(key, h.hashMask)
	if offset == HashTableInvalidOffset {
		return false, fmt.Errorf("invalid param: key(type: %s)", reflect.TypeOf(key).Kind().String())
	}

	h.hash[offset].rwLock.Lock()
	defer h.hash[offset].rwLock.Unlock()

	if h.hash[offset].nodes == nil {
		h.hash[offset].nodes = make([]nodeData, 0)
	}

	IsExisted := false
	nodes := h.hash[offset].nodes
	for i := 0; i < len(nodes); i++ {
		if nodes[i].key == key {
			IsExisted = true
			break
		}
	}

	if !IsExisted {
		h.hash[offset].nodes = append(nodes, *&nodeData{key:key, data:data})
		h.hash[offset].count++
	} else {
			return false, fmt.Errorf("already existed")
	}

	return true, nil
}

func (h *HashTable)Delete(key interface{}) (bool, error) {
	if nil == key {
		return false, fmt.Errorf("invalid param")
	}
	offset := h.getKey(key, h.hashMask)
	if offset == HashTableInvalidOffset {
		return false, fmt.Errorf("invalid param: key(type: %s)", reflect.TypeOf(key).Kind().String())
	}

	hNode := &h.hash[offset]
	hNode.rwLock.Lock()
	defer hNode.rwLock.Unlock()

	if hNode.nodes == nil {
		return true, nil
	}

	nodes := hNode.nodes
	if nodes[0].key == key {
		hNode.nodes = nodes[1:]
		hNode.count--
		return true, nil
	}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].key == key {
			hNode.nodes = append(nodes[0:i-1], nodes[i+1:]...)
			hNode.count--
		}
	}

	return true, nil
}

func (h *HashTable)Search(key interface{}) (interface{}, error) {
	if nil == key {
		return nil, fmt.Errorf("invalid param")
	}
	offset := h.getKey(key, h.hashMask)
	if offset == HashTableInvalidOffset {
		return nil, fmt.Errorf("invalid param: key(type: %s)", reflect.TypeOf(key).Kind().String())
	}

	hNode := h.hash[offset]
	hNode.rwLock.RLock()
	defer hNode.rwLock.RUnlock()

	if h.hash[offset].nodes == nil {
		return nil, fmt.Errorf("not existed")
	}

	nodes := h.hash[offset].nodes
	for i := 0; i < len(nodes); i++ {
		if nodes[i].key == key {
			return nodes[i].data, nil
		}
	}

	return nil, fmt.Errorf("not existed")
}