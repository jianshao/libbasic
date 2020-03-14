package algorithm

import (
	"fmt"
	"github.com/libbasic/datastruct"
)

type Lru struct {
	history *datastruct.DLinkList       /* 历史访问链表 */
	keys map[string]interface{}         /* 节点hash map */
	maxNode int                         /* 节点数上限 */
	currNodeCount int                   /* 当前节点数 */
}

func NewLru() *Lru {
	return &Lru{
		history:datastruct.NewDLinkList(),
		keys:make(map[string]interface{}),
	}
}

func (l *Lru)Save(key string, value interface{}) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("invalid param")
	}

	if l.keys[key] == nil {
		if l.currNodeCount >= l.maxNode {
			l.history.DeleteTail()
		}
		if _, err := l.history.Add2Head(key); err != nil {
			return false, err
		}
		l.keys[key] = value
		l.currNodeCount++
	} else {
		if _, err := l.history.Move2Head(key); err != nil {
			return false, err
		}
		l.keys[key] = value
	}

	return true, nil
}

func (l *Lru)Get(key string) (interface{}, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid param")
	}

	if l.keys[key] == nil {
		return nil, fmt.Errorf("not existed")
	}

	if _, err := l.history.Move2Head(key); err != nil {
		return nil, err
	}

	return l.keys[key], nil
}