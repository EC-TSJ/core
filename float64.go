/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type Float64 float64

// type Float64 struct {
// 	Int int
// }

func (f64 Float64) ToFloat32() float32 {
	return float32(f64)
}

func (f64 Float64) ToFloat64() Float64 {
	return f64
}

// // Representa el más largo valor posible de un Float64.
// func (f64 Float64) Max() int32 {
// 	return 0x7fffffff
// }

// // Representa el más corto valor posible de un Float64.
// func (f64 Float64) Min() int32 {
// 	return -0x80000000
// }

// interface IComparer
func (m Float64) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (f64 Float64) CompareTo(vInt32 int) int {
	if w, ok := T(f64).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m Float64) CompareIf(vInt32 T) int {
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
func (f64 Float64) CompareToIf(vInt32 T) int {
	if w, ok := T(f64).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m Float64) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (f64 Float64) EqualsTo(vInt32 int) bool {
	if w, ok := T(f64).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m Float64) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(Float64)
}

// Compara dos valores. Da true, false.
func (f64 Float64) EqualsToIf(vInt32 T) bool {
	if w, ok := T(f64).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (f64 Float64) GetHashCode() float64 {
	return float64(f64)
}

// Parsea una cadena a entero
func (f64 Float64) Parse(v32 string) float64 {
	s, _ := strconv.ParseInt(v32, 10, 64)
	return float64(s)
}

// Retorna el TypeCode del dato
func (f64 Float64) GetTypeCode() TypeCode {
	return _FLOAT64
}
