package binary

type IBinary interface {
	Compare(a, b interface{}) int
	Len() int
	GetByIndex(pos int) interface{}
}

func search(b IBinary, s, e int, data interface{}) int {
	if b.Compare(b.GetByIndex(s), data) == 0 {
		return s
	}
	if b.Compare(b.GetByIndex(e), data) == 0 {
		return e
	}

	if s >= e {
		return -1
	}

	mid := (s + e) / 2
	result := b.Compare(b.GetByIndex(mid), data)
	if result > 0 {
		return search(b, s, mid - 1, data)
	} else if result < 0 {
		return search(b, mid + 1, e, data)
	} else {
		return mid
	}

}

func Search(b IBinary, data interface{}) int {
	return search(b, 0, b.Len() - 1, data)
}

