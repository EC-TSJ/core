/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type Int int

// type Int struct {
// 	Int int
// }

func (i32 Int) ToInt8() int8 {
	return int8(i32)
}

func (i32 Int) ToInt16() int16 {
	return int16(i32)
}

func (i32 Int) ToInt32() int32 {
	return int32(i32)
}

func (i32 Int) ToInt64() int64 {
	return int64(i32)
}

// Representa el más largo valor posible de un Int.
func (i32 Int) Max() int {
	return 9223372036854775807
}

// Representa el más corto valor posible de un Int.
func (i32 Int) Min() int {
	return -9223372036854775808
}

// interface IComparer
func (m Int) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (i32 Int) CompareTo(vInt32 int) int {
	if w, ok := T(i32).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m Int) CompareIf(vInt32 T) int {
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
func (i32 Int) CompareToIf(vInt32 T) int {
	if w, ok := T(i32).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m Int) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (i32 Int) EqualsTo(vInt32 int) bool {
	if w, ok := T(i32).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m Int) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(Int)
}

// Compara dos valores. Da true, false.
func (i32 Int) EqualsToIf(vInt32 T) bool {
	if w, ok := T(i32).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (i32 Int) GetHashCode() int {
	return int(i32)
}

// Parsea una cadena a entero
func (i32 Int) Parse(v32 string) int {
	s, _ := strconv.ParseInt(v32, 10, 64)
	return int(s)
}

// Retorna el TypeCode del dato
func (i32 Int) TypeCode() TypeCode {
	return _INT32
}
