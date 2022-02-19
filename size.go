package core

// Type of data
type (
	ISize int

	// sizeSt struct
	sizeSt struct {
		// Small
		Small ISize
		// Medium
		Medium ISize
		// Large
		Large ISize
		// EXtraLarge
		ExtraLarge ISize
	}
	Esize = sizeSt
)

const (
	// values
	__small__ ISize = iota + 1
	__medium__
	__large__
	__extra_large__
)

// Enum of object's size
// @return {*Size}
var ESize *Esize = &sizeSt{Small: __small__, Medium: __medium__, Large: __large__, ExtraLarge: __extra_large__}

// Name of object
// @return {string}
// @@interface Stringer
func (s ISize) String() string {
	return [...]string{"Small", "Medium", "Large", "ExtraLarge"}[s-1]
}
