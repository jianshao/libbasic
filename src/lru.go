package libbasic

type Lru struct {
	data interface{}
	history []string
}