package core

import (
	"encoding/base64"
	"strings"
)

// ? @public @enum Decorators
var Decorators = &struct {
	Capitalize, MakeItalic, MakeParagraph, MakeBold, ToLower, ToUpper, ToBase64, ToBase32 *Decorator[func(func(...string) string, ...string) string]
	//ToInterface                                                               *Decorator[func(func(any) any, any) any]
	// ToFloat  *Decorator[func(func(any) float64, any) float64]
	// ToInt    *Decorator[func(func(any) int, any) int]
	// ToString *Decorator[func(func(any) any, any) any]
	//To *Decorator[func(func(int) int, int) int]
}{
	Capitalize:    (&Decorator[func(func(...string) string, ...string) string]{}).New(Capitalize),
	MakeItalic:    (&Decorator[func(func(...string) string, ...string) string]{}).New(MakeItalic),
	MakeParagraph: (&Decorator[func(func(...string) string, ...string) string]{}).New(MakeParagraph),
	MakeBold:      (&Decorator[func(func(...string) string, ...string) string]{}).New(MakeBold),
	//ToInterface:   (&Decorator[func(func(any) any, any) any]{}).New(ToInterface),
	// ToFloat:  (&Decorator[func(func(any) float64, any) float64]{}).New(ToFloat),
	// ToInt:    (&Decorator[func(func(any) int, any) int]{}).New(ToInt),
	// ToString: (&Decorator[func(func(any) any, any) any]{}).New(ToString),
	//To:       (&Decorator[func(func(int) int, int) int]{}).New(To),
	ToLower:  (&Decorator[func(func(...string) string, ...string) string]{}).New(ToLower[func(...string) string]),
	ToUpper:  (&Decorator[func(func(...string) string, ...string) string]{}).New(ToUpper[func(...string) string]),
	ToBase64: (&Decorator[func(func(...string) string, ...string) string]{}).New(ToBase64[func(...string) string]),
	ToBase32: (&Decorator[func(func(...string) string, ...string) string]{}).New(ToBase32[func(...string) string]),
}

/**-----------------------------------------------------------------------------
 *------------------------------------------------------------------------------
 */

/**
 * Capitaliza una cadena
 *
 *! @param {any}
 *! @return {func() any}
 */
// ? @decorator Capitalize(inner, any) any
func Capitalize[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any(strings.Title(any(inner(x[0])).(string))).(T)
}

/**
  * Hace itálica una cadena
  *! @param {any}
	*! @return {func() any}
*/
// ? @decorator MakeItalic(inner, any) any
func MakeItalic[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any("<i>" + any(inner(x[0])).(string) + "</i>").(T)
}

/**
  * Hace un párrafo de una cadena
  *! @param {any}
	*! @return {func() any}
*/
// ? @decorator MakeParagraph(inner, any) any
func MakeParagraph[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any("<p>" + any(inner(x[0])).(string) + "</p>").(T)
}

// // !++
// // ? @MakeItalic(inner, s)
// var v = (&Decorator[func(func(any) any, any) any]{}).New(MakeItalic)

// func functionWithDecorator1(s any) any {
// 	if decorators.MakeItalic.Execute {
// 		decorators.MakeItalic.Set(false)
// 		return decorators.MakeItalic.Decorate(functionWithDecorator1, s)
// 	}
// 	return s
// }

// //!--

// // !++
// // ? @MakeParagraph(inner, s)
// var v2 = (&Decorator[func(func(any) any, any) any]{}).New(MakeParagraph)

// func functionWithDecorator2(s any) any {
// 	if decorators.MakeParagraph.Execute {
// 		decorators.MakeParagraph.Set(false)
// 		return decorators.MakeParagraph.Decorate(functionWithDecorator2, s)
// 	}
// 	return s
// }

//!--
/**
  * Hace negrita de una cadena
  *! @param {any}
	*! @return {func() any}
*/
// ? @decorator MakeBold(inner, any) any
func MakeBold[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any("<b>" + any(inner(x[0])).(string) + "</b>").(T)
}

//-------------------------------------------------------------
// data type

// /**
//   * Hace un interface{} de un dato
//   *! @param {any}
// 	*! @return {func() any}
// */
// // ? @decorator ToInterface(inner, any) any
// func ToInterface(inner func(any) any, x any) any {
// 	return inner(x)
// }

// /* hacerlo con unsafe.Pointer's */
// /**
//   * Hace un float de un dato
//   *! @param {any}
// 	*! @return {func() any}
// */
// // ? @decorator ToFloat(inner, any) any
// func ToFloat[X interface{ func(any) float64 }](inner X, x any) float64 {
// 	return inner(x)
// }

// /**
//   * Hace un int de un dato
//   *! @param {any}
// 	*! @return {func() any}
// */
// // ? @decorator ToInt(inner, any) any
// func ToInt[X interface{ func(any) int }](inner X, x any) int {
// 	return inner(x)
// }

// /**
//   * Hace una any de un dato
//   *! @param {any}
// 	*! @return {func() any}
// */
// // ? @decorator ToString(inner, any) any
// func ToString[X interface{ func(any) any }](inner X, x any) any {
// 	return inner(x)
// }

// /**
//   * Hace un cast de un dato
//   *! @param {any}
// 	*! @return {func() any}
// */
// // ? @decorator To(inner, any) any
// func To[T Ordered, X interface{ func(T) T }](inner X, x T) T {
// 	return inner(x)
// }

//----------------------------------------------------------------
// Minusculas, Mayusculas
/**
  * Convierte una any de un dato, a minúsculas
  *! @param {func(any) any}
	*! @return {func(any) any}
*/
// ? @decorator ToLower(inner, any) any
func ToLower[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any(strings.ToLower(any(inner(x[0])).(string))).(T)
}

/**
  * Convierte una any de un dato, a mayúsculas
  *! @param {func(any) any}
	*! @return {func(any) any}
*/
// ? @decorator ToUpper(inner, any) any
func ToUpper[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any(strings.ToUpper(any(inner(x[0])).(string))).(T)
}

//_______________________________________________________________
// Bases
/**
  * Convierte una any de un dato, a Base64
  *! @param {func(any) any}
	*! @return {func(any) any}
*/
// ? @decorator ToBase64(inner, any) any
func ToBase64[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any(base64.StdEncoding.EncodeToString([]byte(any(inner(x[0])).(string)))).(T)
}

/**
  * Convierte una any de un dato, a Base32
  *! @param {func(any) any}
	*! @return {func(any) any}
*/
// ? @decorator ToBase32(inner, any) any
func ToBase32[X interface{ func(...T) T }, T comparable](inner X, x ...T) T {
	return any(base64.StdEncoding.EncodeToString([]byte(any(inner(x[0])).(string)))).(T)
}

// //-------------------------------------------------------------------
// // Append, Prepend
// /**
//   * Hace un prefacio a una cadena
//   *! @param {any func(any) any func(any) any}
// 	*! @return {func(any) any}
// */
// // ? @decorator AppendDecorator(inner, any) any
// func AppendDecorator[T Ordered, X interface{ func(T) T }](inner X, x, s T) T {
// 	return inner(s + x)
// }

// /**
//   * Hace un epifacio a una cadena
//   *! @param {any,  func(any) any func(any) any}
// 	*! @return {func(any) any}
// */
// // ? @decorator PrependDecorator(inner, any) any
// func PrependDecorator[T Ordered, X interface{ func(T) T }](inner X, x, s T) T {
// 	return inner(x + s)
// }

// func main() {
// 	fmt.Println(functionWithDecorator1("gilipollas"))
// 	fmt.Println(functionWithDecorator2("gilipollas"))
// }
