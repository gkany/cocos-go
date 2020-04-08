package types

const (
	INT = iota
	NUMBER
	STRING
	BOOL
	TABLE
	FUNCTION
)

type LuaObject struct {
	Type  uint64
	Value interface{}
}