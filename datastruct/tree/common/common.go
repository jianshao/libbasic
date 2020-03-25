package common

type treeNode struct {
	l, r *treeNode
	data interface{}
}

type Tree struct {
	root *treeNode
}

func GetNodeCount(t *Tree) int {
	return 0
}

func GetMaxDepth(t *Tree) int {
	return 0
}

func TransferToDLinkList(t *Tree)  {

}

func GetNodeCountInKDepth(t *Tree) int {
	return 0
}

func IsMirror(t1, t2 *Tree) bool {
	return false
}

func IsFullTree(t *Tree) bool {
	return false
}

func IsAvlTree(t *Tree) bool {
	return false
}