/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import "maps"

// *******************************************************************************
// ? @public @class Registry
// *******************************************************************************
type Registry map[string]any

var (
	// GlobalRegistry
	GlobalRegistry Registry = Registry{
		// "HashTable": hashtable.HashTable{},
		// //"CLI":       cli.Command{},
		// "CIF":   codes.CIF{},
		// "NIE":   codes.NIE{},
		// "NIF":   codes.NIF{},
		// "CCC":   codes.CCC{},
		// "NAFSS": codes.NAFSS{},
		// "CARD":  codes.CARD{},
		// "IBAN":  codes.IBAN{},
	}

	//!+
	__registry Registry = make(Registry)
)

// Ancho de GlobalRegistry
// ? core.Method.Normal
// ! @return {int}
func (r Registry) Length() int {
	return len(GlobalRegistry)
}

// Retorna un valor de GlobalRegistry
// ? core.Method.Normal
// ! @return {int}
func (r Registry) Get(s string) any {
	return GlobalRegistry[s]
}

/**
 * Crea un registro único, suma de los registros globales, locales, etc
 * Los valores de los registros de la derecha tienen preferencia sobre los valores de la izquierda, se
 * han de poner primero globales y luego locales.
 * <CODE>
 *?		golang.MakeRegistry(GlobalRegistry, LocalRegistry)
 * </CODE>
 *! @param{...map[string]any} los registros a sumar
 */
func MakeRegistry(m /*...Registry*/ Registry) {
	__registry := maps.Clone(GlobalRegistry)
	maps.Copy(__registry, m)
}

/**
 * Obtiene un valor del registro
 *! @param{string}
 *! @return {any}
 */
func GetRegistry(s string) any {
	return __registry[s]
}

//!-

// /**
//  * Obtiene el GlobalRegistry
//  *! @return {Registry}
//  */
// func GRegistry() Registry {
// 	return GlobalRegistry
// }
