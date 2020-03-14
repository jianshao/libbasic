package test

import (
	"github.com/libbasic/datastruct"
	"testing"
)

func TestTrie_Add(t *testing.T)  {
	trie := datastruct.NewTrie()
	AssertNotNil(t, trie)

	/* 输入参数错误类型 */
	var k []string = nil
	r, err, _ := trie.AddNode(k, 1)
	AssertFalse(t, r)
	AssertEqual(t, err, "")

	k = []string{"a", "", "c"}
	r, err, _ = trie.AddNode(k, 1)
	AssertFalse(t, r)
	AssertEqual(t, err, "")

	r, err, _ = trie.AddNode(k, nil)
	AssertFalse(t, r)
	AssertEqual(t, err, "")

	/* 正常测试例 */
	keys := []string{"a", "b", "n"}
	r, _, _ = trie.AddNode(keys, 1)
	AssertTrue(t, r)

	data, err := trie.Search(keys)
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	/* 重复添加错误 */
	r, err, d := trie.AddNode(keys, 2)
	AssertFalse(t, r)
	AssertEqual(t, err, "")
	AssertEqual(t, d, 1)
}

func TestTrie_Delete(t *testing.T)  {
	trie := datastruct.NewTrie()
	AssertNotNil(t, trie)

	/* 参数输入错误类型 */
	r, err := trie.Delete(nil)
	AssertFalse(t, r)
	AssertEqual(t, err, "")

	k := []string{"a", "", "c"}
	r, err = trie.Delete(k)
	AssertFalse(t, r)
	AssertEqual(t, err, "")

	/* 删除节点不存在 */
	keys := []string{"a", "b", "c"}
	r, err = trie.Delete(keys)
	AssertTrue(t, r)
	AssertNil(t, err)

	/* 正常删除 */
	r, err, d := trie.AddNode(keys, 1)
	AssertTrue(t, r)
	AssertNil(t, err)
	AssertNil(t, d)

	r, err = trie.Delete(keys)
	AssertTrue(t, r)
	AssertNil(t, err)
}

func TestTrie_Search(t *testing.T)  {
	trie := datastruct.NewTrie()
	AssertNotNil(t, trie)

	/* 输入参数错误 */
	d, err := trie.Search(nil)
	AssertNil(t, d)
	AssertEqual(t, err, "")

	k := []string{"", "b", "c"}
	d, err = trie.Search(k)
	AssertNil(t, d)
	AssertEqual(t, err, "")

	/* 节点不存在 */
	keys := []string{"a", "b", "c"}
	d, err = trie.Search(keys)
	AssertNil(t, d)
	AssertEqual(t, err, "")

	/* 节点存在，正常查询 */
	r, err, d := trie.AddNode(keys, 1)
	AssertTrue(t, r)
	AssertNil(t, err)
	AssertNil(t, d)

	d, err = trie.Search(keys)
	AssertEqual(t, d, 1)
	AssertNil(t, err)
}