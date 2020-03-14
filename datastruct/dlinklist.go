package datastruct

import (
	"fmt"
)

type dLinkListNode struct {
	p *dLinkListNode
	n *dLinkListNode
	data interface{}
}

type DLinkList struct {
	head *dLinkListNode
	tail *dLinkListNode
	nodeCount int
}

func NewDLinkList() *DLinkList {
	return &DLinkList{
		head:nil,
		tail:nil,
		nodeCount:0,
	}
}

func (d *DLinkList)Add2Head(data interface{}) (bool, error) {
	if nil == data {
		return false, fmt.Errorf("invalid param")
	}

	n := &dLinkListNode{
		p:nil,
		n:d.head,
		data:data,
	}

	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		d.head.p = n
		d.head = n
	}
	return true, nil
}

func (d *DLinkList)Add2Tail(data interface{}) (bool, error) {
	if nil == data {
		return false, fmt.Errorf("invalid param")
	}

	n := &dLinkListNode{
		p:d.tail,
		n:nil,
		data:data,
	}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		d.tail.n = n
		d.tail = n
	}
	return true, nil
}

func (d *DLinkList)Append(data interface{}) (bool, error) {
	return d.Add2Tail(data)
}

func (d *DLinkList)Delete(data interface{}) (bool, error) {
	if nil == data {
		return false, fmt.Errorf("invalid param")
	}

	if d.head.data == data {
		temp := d.head
		d.head = d.head.n
		d.head.p = nil
		temp.n = nil
		d.nodeCount--
		return true, nil
	}
	if d.tail.data == data {
		temp := d.tail
		d.tail = d.tail.p
		d.tail.n = nil
		temp.p = nil
		d.nodeCount--
		return true, nil
	}

	node := d.head.n
	for ; node != nil; {
		if data == node.data {
			node.p.n = node.n
			node.n.p = node.p
			d.nodeCount--
			break
		}
	}

	return true, nil
}

func (d *DLinkList)DeleteHead() (bool, error) {
	return d.Delete(d.head.data)
}

func (d *DLinkList)DeleteTail() (bool, error) {
	return d.Delete(d.tail.data)
}


func (d *DLinkList)Get(data interface{}) (interface{}, error) {
	if nil == data {
		return false, fmt.Errorf("invalid param")
	}

	head := d.head
	for ; head != nil; {
		if head.data == data {
			return head.data, nil
		}
		head = head.n
	}
	return nil, fmt.Errorf("not existed")
}

func (d *DLinkList)GetHead() (interface{}, error) {
	if d.head == nil {
		return nil, fmt.Errorf("not existed")
	}
	return d.head.data, nil
}

func (d *DLinkList)GetTail() (interface{}, error) {
	if d.tail == nil {
		return nil, fmt.Errorf("not existed")
	}
	return d.tail.data, nil
}



func (d *DLinkList)Move2Head(data interface{}) (bool, error) {
	if _, err := d.Delete(data); err != nil {
		return false, err
	}
	if _, err := d.Add2Head(data); err != nil {
		return false, err
	}
	return true, nil
}

func (d *DLinkList)Move2Tail(data interface{}) (bool, error) {
	if _, err := d.Delete(data); err != nil {
		return false, err
	}
	if _, err := d.Add2Tail(data); err != nil {
		return false, err
	}
	return true, nil
}

func (d *DLinkList)GetNodeCount() int {
	return d.nodeCount
}