package hashtable

import "fmt"

type IHashTable interface {
	GetOffSetMask() int
	GetKeyFromNode(node interface{}) int

	GetLenInList(list interface{}) int
	GetListByKey(key int) interface{}
	Add2List(list interface{}, node interface{}) (bool, error)
	DeleteFromList(list interface{}, node interface{})
	GetNodeByIndex(list interface{}, pos int) interface{}

	CopyNode(src interface{}) interface{}
}

func Add(i IHashTable, data interface{}) (bool, error) {
	key := i.GetKeyFromNode(data)
	list := i.GetListByKey(key)
	return i.Add2List(list, data)
}

func Delete(i IHashTable, node interface{})  {
	key := i.GetKeyFromNode(node)
	list := i.GetListByKey(key)
	i.DeleteFromList(list, node)
}

func transfer(src, dst IHashTable, copy bool) (bool, error) {
	srcMask := src.GetOffSetMask()

	for i := 0; i < srcMask; i++ {
		srcList := src.GetListByKey(i)
		l := src.GetLenInList(srcList)
		for nodeIndex := 0; nodeIndex < l; nodeIndex++ {
			node := src.GetNodeByIndex(srcList, nodeIndex)
			if copy {
				node = dst.CopyNode(node)
			}
			if _, err := Add(dst, node); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func Copy(src, dst IHashTable) (bool, error) {
	return transfer(src, dst, true)
}

func ReHash(src, dst IHashTable) (bool, error) {
	return transfer(src, dst, false)
}

func Search(src IHashTable, node interface{}) (interface{}, error) {
	key := src.GetKeyFromNode(node)
	list := src.GetListByKey(key)
	l := src.GetLenInList(list)
	for i := 0; i < l; i++ {
		node := src.GetNodeByIndex(list, i)
		if src.GetKeyFromNode(node) == key {
			return node, nil
		}
	}

	return nil, fmt.Errorf("not existed")
}