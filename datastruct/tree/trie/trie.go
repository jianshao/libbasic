package trie


import (
	"fmt"
)

type trieNode struct {
	data interface{}                /* 用户数据 */
	children map[string]*trieNode   /* 子节点 */
	bIsLeafNode bool                /* 叶子节点标识，只有该节点是叶子节点时才能使用用户数据 */
}


type TrieRoot struct {
	root *trieNode
}

func NewTrie() *TrieRoot {
	t := &TrieRoot{
		root:new(trieNode),
	}
	t.root.children = make(map[string]*trieNode)

	return t
}

func (t *TrieRoot) AddNode(key []string, data interface{}) (bool, error, interface{}) {
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] != nil {
			node = node.children[key[i]]
		} else {
			node.children[key[i]] = &trieNode{
				data: nil,
				children: make(map[string]*trieNode),
				bIsLeafNode: false,
			}

			node = node.children[key[i]]
		}
	}

	if node.bIsLeafNode {
		return false, fmt.Errorf("already existed"), node.data
	} else {
		node.bIsLeafNode = true
		node.data = data
	}
	return true, nil, nil
}

func (t *TrieRoot) Delete(key []string) (bool, error) {
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] != nil {
			node = node.children[key[i]]
		} else {
			/* 如果中间有某个节点不存在，则要删除的节点一定不存在 */
			return false, nil
		}
	}
	/* 只需要将叶子节点标识去掉，该节点就不会被当作叶子节点使用，内容也不会被访问 */
	node.bIsLeafNode = false
	return true, nil
}

func (t *TrieRoot) Update(key []string, newData interface{}) (bool, error) {
	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] != nil {
			node = node.children[key[i]]
		} else {
			return false, fmt.Errorf("not existed")
		}
	}

	if node.bIsLeafNode {
		node.data = newData
	} else {
		return false, fmt.Errorf("not a leaf node")
	}
	return true, nil
}

func (t *TrieRoot) Search(key []string) (interface{}, error) {
	var data interface{} = nil

	node := t.root
	for i := 0; i < len(key); i++ {
		if node.children[key[i]] != nil {
			node = node.children[key[i]]
		} else {
			return nil, fmt.Errorf("not existed")
		}
	}

	if node.bIsLeafNode {
		data = node.data
	} else {
		return nil, fmt.Errorf("not leaf node")
	}
	return data, nil
}

func (t *TrieRoot)DeleteTrie() (bool, error) {

	return true, nil
}
