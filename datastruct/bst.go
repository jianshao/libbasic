package datastruct

import "fmt"

type compare func(a interface{}, b interface{}) (int)

type bstNode struct{
	p *bstNode
	l *bstNode
	r *bstNode
	data interface{}
}

type Bst struct{
	root *bstNode
	cmp compare
	nodeCount int
}

func NewBst(cmp compare) *Bst {
	return &Bst{
		root:nil,
		cmp: cmp,
		nodeCount: 0,
	}
}

func newBstNode(p *bstNode, data interface{}) *bstNode {
	return &bstNode{
		p:p,
		l:nil,
		r:nil,
		data:data,
	}
}

func (b *Bst)Add(data interface{}) (bool, error) {
	if b.root == nil {
		b.root = newBstNode(nil, data)
		return true, nil
	}

	root := b.root
	for {
		r := b.cmp(root.data, data)
		if r > 0 {
			if root.r == nil {
				root.r = newBstNode(root, data)
				break
			} else {
				root = root.r
			}
		} else {
			if root.l == nil {
				root.l = newBstNode(root, data)
				break
			} else {
				root = root.l
			}
		}
	}
	return true, nil
}

func findBiggestNode(root *bstNode) (*bstNode, error) {
	if root == nil {
		return nil, nil
	}
	rNode := root
	for {
		if rNode.r != nil {
			rNode = rNode.r
		} else {
			break
		}
	}
	return rNode, nil
}

func (b *Bst)Delete(data interface{}) {
	node, err := b.findTheNode(data)
	if err != nil {
		return
	}
	newNode, err := findBiggestNode(node.l)
	if newNode != nil {
		/* 在左子树中能找到最大节点，将最大节点更新为根节点，同时更新左子树 */
		if newNode == node.l {
			/* 左子树的根节点就是最大节点，此时不需要任何更新 */
		} else {
			/* 先将最大节点从左子树中删除，然后将左子树更新为最大节点的左节点 */
			if newNode.l != nil {
				newNode.p.r = newNode.l
			} else {
				newNode.p.r = nil
			}
			newNode.l = node.l
		}
		/* 将右子树更新为最大节点的右子树 */
		newNode.r = node.r
	} else {
		/* 不存在左子树时，直接将右子树提上来 */
		newNode = node.r
	}

	/* 更新父节点 */
	if node == b.root {
		b.root = newNode
	} else {
		if b.cmp(node.p.data, node.data) > 0 {
			node.p.l = newNode
		} else {
			node.p.r = newNode
		}
	}
}

func (b *Bst)findTheNode(data interface{}) (*bstNode, error) {

	for root := b.root; root != nil; {

		r := b.cmp(root.data, data)
		if r == 0 {
			return root, nil
		} else if r < 0 {
			root = root.r
		} else {
			root = root.l
		}
	}
	return nil, fmt.Errorf("not existed")
}

func (b *Bst)IsExist(data interface{}) bool {
	if _, err := b.findTheNode(data); err != nil {
		return false
	}

	return true
}

func mid(root *bstNode, r []interface{}) {
	if root == nil {
		return
	}

	mid(root.l, r)
	r = append(r, root.data)
	mid(root.r, r)
}

func (b *Bst)GetSortedList() ([]interface{}, error) {
	r := make([]interface{}, b.nodeCount)
	mid(b.root, r)
	return r, nil
}

