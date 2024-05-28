/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/
/* Equipación para usar decoradores. */
/* GlobalDecorators tiene decoradores usados globalmente */

package core

import "maps"

// una manera de definir decoradores, mediante un map
// Decorator[T] es la definición de los decoradores.
// Decorators es una @enum que contiene las definiciones de decoradores
// GlobalDecorators es un map que contiene decoradores

// *******************************************************************************
// ? @public @class Decoras[T]
// *******************************************************************************
type Decoras[T comparable] map[string]func(func(...T) T, ...T) T

var (
	// GlobalDecorators
	GlobalDecorators = Decoras[string]{
		"Capitalize":    Capitalize[func(...string) string],
		"MakeItalic":    MakeItalic[func(...string) string],
		"MakeParagraph": MakeParagraph[func(...string) string],
		"MakeBold":      MakeBold[func(...string) string],
		"ToLower":       ToLower[func(...string) string],
		"ToUpper":       ToUpper[func(...string) string],
		"ToBase64":      ToBase64[func(...string) string],
		"ToBase32":      ToBase32[func(...string) string],
	}

	//!+
	__decorators Decoras[string] = make(Decoras[string])
)

// Ancho de GlobalDecorators
// ? core.Method.Normal
// ! @return {int}
func (s Decoras[T]) Length() int {
	return len(GlobalDecorators)
}

// Retorna un valor de GlobalDecorators
// ? core.Method.Normal
// ! @return {int}
func (s Decoras[T]) Get(ut string) func(func(...T) T, ...T) T {
	return any(GlobalDecorators[ut]).(func(func(...T) T, ...T) T)
}

/**
 * Crea un registro único, suma de los registros globales, locales, etc
 * Los valores de los registros de la derecha tienen preferencia sobre los valores de la izquierda, se
 * han de poner primero globales y luego locales.
 * <CODE>
 *?		golang.MakeSettings(GlobalDecorators, LocalDecorators)
 * </CODE>
 *! @param{...map[string]string} los registros a sumar
 */
func MakeDecoras(m /*...Decoras*/ Decoras[string]) {
	__decorators := maps.Clone(GlobalDecorators)
	maps.Copy(__decorators, m)
}

/**
 * Obtiene un valor del registro de configuraciones
 *! @param{string}
 *! @return {any}
 */
func GetDecora[T comparable](s string) func(func(...T) T, ...T) T {
	return any(__decorators[s]).(func(func(...T) T, ...T) T)
}

//!-

// /**
//  * Obtiene el GlobalDecorators
//  *! @return {Settings}
//  */
// func GDecoras() Decoras {
// 	return GlobalDecorators
// }
