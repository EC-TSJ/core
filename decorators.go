package core

import (
	"encoding/base64"
	"internal/errors"
	"strings"
)

// Función title de strings
//
// ! @param {string}
// ! @return {string}
func title(s string) string {
	sb, n := strings.Split(s, Lit.Space), 0 // separa las palabras
	s = Lit.NullString
	for n < len(sb) {
		ss := []byte{} // string => bytes
		if ss = []byte(sb[n]); ss[0] >= 97 && ss[0] <= 122 {
			ss[0] &= 0b11011111 // resta 32
		}
		// añade palabra a frase final
		s += If[string](n == 0)(string(ss), Lit.Space+string(ss))
		n++ // próxima palabra
	}
	return s
}

// ? @public @enum Decorators
type Decorators[T Ordered] struct {
	// AbstracticusError                                                                     *Decorator[func(func(...any) T) func(ClassType, Object_Of) T]
	// AbstracticusError1                                                                    *Decorator[func(func() T) func(ClassType, Object_Of) T]
	// AbstracticusError2                                                                    *Decorator[func(func(T)) func(ClassType, Object_Of) T]
	// SealedticusError                                                                      *Decorator[func(func(...any) T) func(ClassType) T]
	AbstractError                                                                         *Decorator[func(func(...any) T) func(ClassType, Object_Of) T]
	SealedError                                                                           *Decorator[func(func(...any) T) func(ClassType) T]
	Capitalize, MakeItalic, MakeParagraph, MakeBold, ToLower, ToUpper, ToBase64, ToBase32 *Decorator[func(func(string) string) func(string) string]
}

func (dec *Decorators[T]) New() *Decorators[T] {
	return &Decorators[T]{
		AbstractError: (&Decorator[func(func(...any) T) func(ClassType, Object_Of) T]{}).New(AbstractError),
		SealedError:   (&Decorator[func(func(...any) T) func(ClassType) T]{}).New(SealedError),
		// AbstracticusError:     (&Decorator[func(func(...any) T) func(...any) T]{}).New(AbstractError),
		// SealedticusError:       (&Decorator[func(func(...any) T) func(...any) T]{}).New(SealedError),
		Capitalize:    (&Decorator[func(func(string) string) func(string) string]{}).New(Capitalize),
		MakeItalic:    (&Decorator[func(func(string) string) func(string) string]{}).New(MakeItalic),
		MakeParagraph: (&Decorator[func(func(string) string) func(string) string]{}).New(MakeParagraph),
		MakeBold:      (&Decorator[func(func(string) string) func(string) string]{}).New(MakeBold),
		ToLower:       (&Decorator[func(func(string) string) func(string) string]{}).New(ToLower),
		ToUpper:       (&Decorator[func(func(string) string) func(string) string]{}).New(ToUpper),
		ToBase64:      (&Decorator[func(func(string) string) func(string) string]{}).New(ToBase64),
		ToBase32:      (&Decorator[func(func(string) string) func(string) string]{}).New(ToBase32),
	}
}

/*******************************************************************************
 * Decorador para denotar un error de clase abstracta y colgar el código
 *
 *! @type {X interface { func(...T) T}}
 *! @type {T Number}
 *! @param {X}
 *! @param {ClassType} parms[0]
 *! @param {ClassMethods} parms[1]
 *! @return {T}
 */
func AbstractError[X interface{ func(...any) T }, T Ordered](inner X) func(ClassType, Object_Of) T {
	return func(this0 ClassType, this1 Object_Of) T {
		if this0 == Class.Abstract {
			if this1 == ObjectOf.Class {
				Exception(nil)(func() { Throw(errors.ErrAbstractClass) })
			} else {
				Exception(nil)(func() { Throw(errors.ErrMethodAbstractClass) })
			}
		}
		var (
			anyWhere []any
		)
		anyWhere = append(anyWhere, this0, this1)
		return inner(anyWhere)
	}
}

/*******************************************************************************
 * Decorador para denotar un error de clase sealed y colgar el código
 *
 *! @type {X interface { func(...T) T}}
 *! @type {T Number}
 *! @param {X}
 *! @param {ClassType} parms[0]
 *! @param {ClassMethods} parms[1]
 *! @return {T}
 */
func SealedError[X interface{ func(...any) T }, T Ordered](inner X) func(ClassType) T {
	return func(this0 ClassType) T {
		if this0 == Class.Sealed {
			Exception(nil)(func() { Throw(errors.ErrSealedClass) })
		}
		return inner(this0)
	}
}

// /*******************************************************************************
//  * Decorador para denotar un error de clase abstracta y colgar el código
//  *
//  *! @type {X interface { func(...T) T}}
//  *! @type {T Number}
//  *! @param {X}
//  *! @param {ClassType} parms[0]
//  *! @param {ClassMethods} parms[1]
//  *! @return {T}
//  */
// func AbstracticusError[X interface{ func(...any) T }, T Ordered](inner X) func(parms ...any) T {
// 	return func(this ...any) T {
// 		if this[0].(ClassType) == Class.Abstract {
// 			if this[1].(Object_Of) == ObjectOf.Class {
// 				Exception(nil)(func() { Throw(errors.ErrAbstractClass) })
// 			} else {
// 				Exception(nil)(func() { Throw(errors.ErrMethodAbstractClass) })
// 			}
// 		}
// 		return inner(this)
// 	}
// }

// /*******************************************************************************
//  * Decorador para denotar un error de clase sealed y colgar el código
//  *
//  *! @type {X interface { func(...T) T}}
//  *! @type {T Number}
//  *! @param {X}
//  *! @param {ClassType} parms[0]
//  *! @param {ClassMethods} parms[1]
//  *! @return {T}
//  */
// func SealedticusError[X interface{ func(...any) T }, T Ordered](inner X) func(parms ...any) T {
// 	return func(this ...any) T {
// 		if this[0].(ClassType) == Class.Sealed {
// 			Exception(nil)(func() { Throw(errors.ErrSealedClass) })
// 		}
// 		return inner(this)
// 	}
// }

/**
 * Capitaliza una cadena
 *
 *! @param {any} fn a llamar
 *! @return {func() any} parms para fn a llamar
 */
// ? @decorator Capitalize(inner, any) any
func Capitalize[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return title(inner(parms))
	}
}

/**
  * Hace itálica una cadena
  *! @param {any} fn a llamar
	*! @return {func() any} parms para fn a llamar
*/
// ? @decorator MakeItalic(inner, any) any
func MakeItalic[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return "<i>" + inner(parms) + "</i>"
	}
}

/**
  * Hace un párrafo de una cadena
  *! @param {any} fn a llamar
	*! @return {func() any} parms para fn a llamar
*/
// ? @decorator MakeParagraph(inner, any) any
func MakeParagraph[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return "<p>" + inner(parms) + "</p>"
	}
}

//!--
/**
  * Hace negrita de una cadena
  *! @param {any} fn a llamar
	*! @return {func() any} parms para fn a llamar
*/
// ? @decorator MakeBold(inner, any) any
func MakeBold[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return "<b>" + inner(parms) + "</b>"
	}
}

// Minusculas, Mayusculas
/**
  * Convierte una any de un dato, a minúsculas
  *! @param {func(any) any} fn a llamar
	*! @return {func(any) any} parms para fn a llamar
*/
// ? @decorator ToLower(inner, any) any
func ToLower[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return strings.ToLower(inner(parms))
	}
}

/**
  * Convierte una any de un dato, a mayúsculas
  *! @param {func(any) any} fn a llamar
	*! @return {func(any) any} parms para fn a llamar
*/
// ? @decorator ToUpper(inner, any) any
func ToUpper[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return strings.ToUpper(inner(parms))
	}
}

//_______________________________________________________________
// Bases
/**
  * Convierte una any de un dato, a Base64
  *! @param {func(any) any} fn a llamar
	*! @return {func(any) any} parms para fn a llamar
*/
// ? @decorator ToBase64(inner, any) any
func ToBase64[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return base64.StdEncoding.EncodeToString([]byte(inner(parms)))
	}
}

/**
  * Convierte una any de un dato, a Base32
  *! @param {func(any) any} fn a llamar
	*! @return {func(any) any} parms para fn a llamar
*/
// ? @decorator ToBase32(inner, any) any
func ToBase32[X interface{ func(string) string }](inner X) func(parms string) string {
	return func(parms string) string {
		return base64.StdEncoding.EncodeToString([]byte(inner(parms)))
	}
}
