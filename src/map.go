package libbasic

type Map struct{
	dataType string
	root libbasic.TrieRoot
}

func NewMap(dataType string) *Map {
	return &Map{
		dataType:dataType,
		root: libbasic.NewTrieRoot(),
	}
}

func (m *Map)Add(key string, data interface{}) (bool, error) {
	return true, nil
}

func (m *Map)Delete(key string) {
	return 
}

func (m *Map)Search(key string) (interface{}, error) {
	return nil, nil
}
