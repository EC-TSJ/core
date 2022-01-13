/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "strconv"

type Float32 float32

// type Float32 struct {
// 	Int int
// }

func (f32 Float32) ToFloat32() float32 {
	return float32(f32)
}

func (f32 Float32) ToFloat64() float64 {
	return float64(f32)
}

// // Representa el más largo valor posible de un Float32.
// func (f32 Float32) Max() int32 {
// 	return 0x7fffffff
// 	//0xffffffffffffffffL;
// 	//0x7fffffffffffffffL;
// 	//0x7FFF;
// 	//0x7F;
// }

// // Representa el más corto valor posible de un Float32.
// func (f32 Float32) Min() int32 {
// 	return -0x80000000
// 	//0x8000000000000000L);
// 	//0x8000
// 	//0x80
// }

// interface IComparer
func (m Float32) Compare(z int) int {
	if int(m) < z {
		return -1
	} else if int(m) > z {
		return 1
	}
	return 0
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (f32 Float32) CompareTo(vInt32 int) int {
	if w, ok := T(f32).(IComparer); ok {
		return w.Compare(vInt32)
	}

	return -1551
}

// interface IComparerIf
func (m Float32) CompareIf(vInt32 T) int {
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
func (f32 Float32) CompareToIf(vInt32 T) int {
	if w, ok := T(f32).(IComparerIf); ok {
		return w.Compare(vInt32)
	}
	return -1551
}

// interface IEquater
func (m Float32) Equals(vInt32 int) bool {
	return int(m) == vInt32
}

// Compara dos valores. Da true, false.
func (f32 Float32) EqualsTo(vInt32 int) bool {
	if w, ok := T(f32).(IEquater); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (m Float32) EqualsIf(vInt32 T) bool {
	if vInt32 == nil {
		return false
	}
	if v, ok := vInt32.(int); ok {
		return v == int(m)
	}
	return m == vInt32.(Float32)
}

// Compara dos valores. Da true, false.
func (f32 Float32) EqualsToIf(vInt32 T) bool {
	if w, ok := T(f32).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}

// Da el Hash de la cantidad.
func (f32 Float32) GetHashCode() float32 {
	return float32(f32)
}

// Parsea una cadena a entero
func (f32 Float32) Parse(v32 string) float32 {
	s, _ := strconv.ParseInt(v32, 10, 32)
	return float32(s)
}

// Retorna el TypeCode del dato
func (f32 Float32) GetTypeCode() TypeCode {
	return _FLOAT32
}
