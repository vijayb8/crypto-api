package http

// ContextKey represents custom key for putting into context
type ContextKey string

// String returns string representation of context key
func (c ContextKey) String() string {
	return string(c)
}
