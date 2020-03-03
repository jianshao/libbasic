package libbasic

type compare func(a interface{}, b interface{}) (int)

type bstNode struct{
	p *bstNode
	l *bstNode
	r *bstNode
	data interface{}
}

type Bst struct{
	root bstNode
	cmp compare
	nodeCount int
}

func NewBst(cmp compare) *Bst {
	return &Bst{
		root:&bstNode{nil, nil, nil},
		cmp: cmp,
		nodeCount: 0,
	}
}

func newBstNode(p *bstNode, data interface{}) (*bstNode) {
	return &bstNode{
		p:p,
		l:nil,
		r:nil,
		data:data,
	}
}

func (b *Bst)Add(data interface{}) (bool, error) {
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

func findBiggestNode(root bstNode) (*bstNode, error) {
	if root == nil {
		return nil, nil
	}
	if root.r == nil {
		if root.l == nil {
			root.p.r = nil
			return root, nil
		} else {
			root.p.r = root.l
			return root, nil
		}
	} else {
		findBiggestNode(root.r)
	}
	return nil, nil
}

func (b *Bst)Delete(data interface{}) {
	node, _ := b.findTheNode(data)
	newNode := findBiggestNode(node)
	if newNode != nil {
		newNode.l = node.l
		newNode.r = node.r
		if b.cmp(node.p.l.data, node.data) == 0 {
			node.p.l = newNode
		} else {
			node.p.r = newNode
		}
	}
}

func (b *Bst)findTheNode(data interface{}) (*bstNode, error) {
	root := b.root
	for {
		if root == nil {
			break
		}

		r := b.cmp(root.data, data)
		if r == 0 {
			return root, nil
		} else if r > 0 {
			root = root.r
		} else {
			root = root.l
		}
	}
	return nil, nil
}

func (b *Bst)IsExist(data interface{}) (bool, error) {
	n, err := b.findTheNode(data)
	if n == nil {
		return false, nil
	}

	return true, nil
}

func mid(root bstNode, r []interface{}) {
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

