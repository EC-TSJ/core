/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type Int64 int64

// type Int64 struct {
// 	Int int
// }

func (i32 Int64) ToInt8() int8 {
	return int8(i32)
}

func (i32 Int64) ToInt16() int16 {
	return int16(i32)
}

func (i32 Int64) ToInt32() int32 {
	return int32(i32)
}

func (i32 Int64) ToInt64() int64 {
	return int64(i32)
}

// uint8       unsigned  8-bit integers (0 to 255)
// uint16      unsigned 16-bit integers (0 to 65535)
// uint32      unsigned 32-bit integers (0 to 4294967295)
// uint64      unsigned 64-bit integers (0 to 18446744073709551615)
// int8        signed  8-bit integers (-128 to 127)
// int16       signed 16-bit integers (-32768 to 32767)
// int32       signed 32-bit integers (-2147483648 to 2147483647)
// int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// Representa el más largo valor posible de un Int64.
func (i32 Int64) Max() int64 {
	return 9223372036854775807
}

// Representa el más corto valor posible de un Int64.
func (i32 Int64) Min() int64 {
	return -9223372036854775808
}

// interface IComparer
func (m Int64) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (i32 Int64) CompareTo(vInt32 int) int {
	if w, ok := T(i32).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m Int64) CompareIf(vInt32 T) int {
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
func (i32 Int64) CompareToIf(vInt32 T) int {
	if w, ok := T(i32).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m Int64) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (i32 Int64) EqualsTo(vInt32 int) bool {
	if w, ok := T(i32).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m Int64) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(Int64)
}

// Compara dos valores. Da true, false.
func (i32 Int64) EqualsToIf(vInt32 T) bool {
	if w, ok := T(i32).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (i32 Int64) GetHashCode() int64 {
	return int64(i32)
}

// Parsea una cadena a entero
func (i32 Int64) Parse(v32 string) int64 {
	s, _ := strconv.ParseInt(v32, 10, 64)
	return int64(s)
}

// Retorna el TypeCode del dato
func (i32 Int64) TypeCode() TypeCode {
	return _INT64
}
