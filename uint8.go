/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type UInt8 uint8

// type UInt8 struct {
// 	Int int
// }

func (i32 UInt8) ToUInt8() uint8 {
	return uint8(i32)
}

func (i32 UInt8) ToUInt16() uint16 {
	return uint16(i32)
}

func (i32 UInt8) ToUInt32() uint32 {
	return uint32(i32)
}

func (i32 UInt8) ToUInt64() uint64 {
	return uint64(i32)
}

// Representa el más largo valor posible de un UInt8.
func (i32 UInt8) Max() uint8 {
	return 255
}

// Representa el más corto valor posible de un UInt8.
func (i32 UInt8) Min() uint8 {
	return 0
}

// interface IComparer
func (m UInt8) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (i32 UInt8) CompareTo(vInt32 int) int {
	if w, ok := T(i32).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m UInt8) CompareIf(vInt32 T) int {
	// compara por defecto
	if vInt32 == nil {
		return 1
	}
	v, ok := vInt32.(int)
	if v, ok = vInt32.(int); ok {
		if int(m) < v {
			return -1
		}
		if int(m) > v {
			return 1
		}
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (i32 UInt8) CompareToIf(vInt32 T) int {
	if w, ok := T(i32).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m UInt8) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (i32 UInt8) EqualsTo(vInt32 int) bool {
	if w, ok := T(i32).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m UInt8) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(UInt8)
}

// Compara dos valores. Da true, false.
func (i32 UInt8) EqualsToIf(vInt32 T) bool {
	if w, ok := T(i32).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (i32 UInt8) GetHashCode() uint8 {
	return uint8(i32)
}

// Parsea una cadena a entero
func (i32 UInt8) Parse(v32 string) uint8 {
	s, _ := strconv.ParseInt(v32, 10, 8)
	return uint8(s)
}

// Retorna el TypeCode del dato
func (i32 UInt8) TypeCode() TypeCode {
	return _INT32
}
