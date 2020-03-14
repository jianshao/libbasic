package algorithm

import "fmt"

type Lru struct {
	history *DLinkList
	keys map[string]interface{}
	maxNode int
	currNodeCount int
}

func NewLru() (*Lru) {
	return &Lru{
		history:NewDLinkList(),
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