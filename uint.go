/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type UInt uint

// type UInt struct {
// 	UInt int
// }

func (i32 UInt) ToUInt8() int8 {
	return int8(i32)
}

func (i32 UInt) ToUInt16() int16 {
	return int16(i32)
}

func (i32 UInt) ToUInt32() int32 {
	return int32(i32)
}

func (i32 UInt) ToUInt64() int64 {
	return int64(i32)
}

// Representa el más largo valor posible de un UInt.
func (i32 UInt) Max() uint {
	return 18446744073709551615
}

// Representa el más corto valor posible de un UInt.
func (i32 UInt) Min() uint {
	return 0
}

// interface IComparer
func (m UInt) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (i32 UInt) CompareTo(vInt32 int) int {
	if w, ok := T(i32).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m UInt) CompareIf(vInt32 T) int {
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
func (i32 UInt) CompareToIf(vInt32 T) int {
	if w, ok := T(i32).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m UInt) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (i32 UInt) EqualsTo(vInt32 int) bool {
	if w, ok := T(i32).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m UInt) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(UInt)
}

// Compara dos valores. Da true, false.
func (i32 UInt) EqualsToIf(vInt32 T) bool {
	if w, ok := T(i32).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (i32 UInt) GetHashCode() uint {
	return uint(i32)
}

// Parsea una cadena a entero
func (i32 UInt) Parse(v32 string) uint {
	s, _ := strconv.ParseInt(v32, 10, 64)
	return uint(s)
}

// Retorna el TypeCode del dato
func (i32 UInt) TypeCode() TypeCode {
	return _INT32
}
