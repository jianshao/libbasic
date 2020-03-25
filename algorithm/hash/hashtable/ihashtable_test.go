package hashtable

import (
	"testing"
)

type ihashNode struct {
	key int
	data interface{}
}

type hashList struct {
	len int
	nodes []ihashNode
}

type hasht struct {
	mask int
	lists []hashList
}

func (h *hasht)GetOffSetMask() int {
	return h.mask
}

func (h *hasht)GetKeyFromNode(node interface{}) int {
	i := node.(ihashNode)
	return i.key
}

func (h *hasht)GetLenInList(list interface{}) int {
	return list.(hashList).len
}

func (h *hasht)GetListByKey(key int) interface{} {
	pos := key % h.mask
	print(pos)
	return &h.lists[pos]
}

func (h *hasht)Add2List(list interface{}, node interface{}) (bool, error) {
	hList := list.(*hashList)
	hList.nodes = append(hList.nodes, node.(ihashNode))
	hList.len++
	return true, nil
}

func (h *hasht)DeleteFromList(list interface{}, node interface{})  {
	hList := list.(hashList)
	for i := 0; i < hList.len; i++ {
		if hList.nodes[i].key == node.(ihashNode).key {
			hList.nodes = append(hList.nodes[0:i-1], hList.nodes[i+1:]...)
			hList.len--
			break
		}
	}

}

func (h *hasht)GetNodeByIndex(list interface{}, pos int) interface{} {
	return list.(hashList).nodes[pos]
}

func (h *hasht)CopyNode(src interface{}) interface{} {
	srcNode := src.(ihashNode)
	newNode := new(ihashNode)
	newNode.key = srcNode.key
	newNode.data = srcNode.data

	return newNode
}

func (h hasht)Init(mask int)  {
	h.mask = mask
	h.lists = make([]hashList, mask)
	for i := 0; i < mask; i++ {
		h.lists[i].len = 0
		h.lists[i].nodes = make([]ihashNode, 0)
	}
}

func (h *hasht)SearchAll(t *testing.T)  {
	t.Helper()

	for i := 0; i < h.mask; i++ {
		r := false
		list := h.lists[i]
		for j := 0; j < list.len; j++ {

			t.Error("-> ", list.nodes[j])
			r = true
		}
		if r {
			t.Error("\n")
		}
	}
}

func TestReHash(t *testing.T) {
	var ht = &hasht{}
	ht.Init(10)

	var node = ihashNode{12, 12}
	Add(ht, node)
	t.Error(ht)
	ht.SearchAll(t)
}