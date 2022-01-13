package core

// Type of data
type iSize int

const (
	// values
	__small__ iSize = iota + 1
	__medium__
	__large__
	__extra_large__
)

// sizeSt struct
type sizeSt struct {
	Small      iSize
	Medium     iSize
	Large      iSize
	ExtraLarge iSize
}

// Enum of object's size
// @return {*Size}
func Size() *sizeSt {
	return &sizeSt{Small: __small__, Medium: __medium__, Large: __large__,
		ExtraLarge: __extra_large__}
}

// Name of object
// @return {string}
// @@interface Stringer
func (s iSize) String() string {
	return [...]string{"Small", "Medium", "Large", "ExtraLarge"}[s-1]
}
