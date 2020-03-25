package trie

import (
	"github.com/libbasic/test"
	"testing"
)

func TestTrie_Add(t *testing.T)  {
	trie := trie2.NewTrie()
	test.AssertNotNil(t, trie)

	/* 输入参数错误类型 */
	var k []string = nil
	r, err, _ := trie.AddNode(k, 1)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")

	k = []string{"a", "", "c"}
	r, err, _ = trie.AddNode(k, 1)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")

	r, err, _ = trie.AddNode(k, nil)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")

	/* 正常测试例 */
	keys := []string{"a", "b", "n"}
	r, _, _ = trie.AddNode(keys, 1)
	test.AssertTrue(t, r)

	data, err := trie.Search(keys)
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	/* 重复添加错误 */
	r, err, d := trie.AddNode(keys, 2)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")
	test.AssertEqual(t, d, 1)
}

func TestTrie_Delete(t *testing.T)  {
	trie := trie2.NewTrie()
	test.AssertNotNil(t, trie)

	/* 参数输入错误类型 */
	r, err := trie.Delete(nil)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")

	k := []string{"a", "", "c"}
	r, err = trie.Delete(k)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "")

	/* 删除节点不存在 */
	keys := []string{"a", "b", "c"}
	r, err = trie.Delete(keys)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	/* 正常删除 */
	r, err, d := trie.AddNode(keys, 1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)
	test.AssertNil(t, d)

	r, err = trie.Delete(keys)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)
}

func TestTrie_Search(t *testing.T)  {
	trie := trie2.NewTrie()
	test.AssertNotNil(t, trie)

	/* 输入参数错误 */
	d, err := trie.Search(nil)
	test.AssertNil(t, d)
	test.AssertEqual(t, err, "")

	k := []string{"", "b", "c"}
	d, err = trie.Search(k)
	test.AssertNil(t, d)
	test.AssertEqual(t, err.Error(), "")

	/* 节点不存在 */
	keys := []string{"a", "b", "c"}
	d, err = trie.Search(keys)
	test.AssertNil(t, d)
	test.AssertEqual(t, err.Error(), "")

	/* 节点存在，正常查询 */
	r, err, d := trie.AddNode(keys, 1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)
	test.AssertNil(t, d)

	d, err = trie.Search(keys)
	test.AssertEqual(t, d, 1)
	test.AssertNil(t, err)
}