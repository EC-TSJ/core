/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"strings"
)

type String string

/**
 * Convierte a minúsculas la cadena
 * @return {string}
 */
func (str String) ToLower() string {
	return strings.ToLower(string(str))
}

/**
 * Convierte a mayúsculas la cadena
 * @return {string}
 */
func (str String) ToUpper() string {
	return strings.ToUpper(string(str))
}

/**
 * Nos dice si la cadena contiene a 'search'
 * @param {string}
 * @return {bool}
 */
func (str String) Contains(search string) bool {
	return strings.Contains(string(str), search)
}

/**
 * Nos dice el índice en donde se encuentra en la cadena
 * @param {string}
 * @return {int}
 */
func (str String) Index(search string) int {
	return strings.Index(string(str), search)
}

/**
 * Divide la cadena en una lista de cadenas divididas por el separador sep
 * @param {string}
 * @return {bool}
 */
func (str String) Split(sep string) []string {
	return strings.Split(string(str), sep)
}

/**
 * Elimina los 'set' caracteres de alrededor de la cadena
 * @param {string}
 * @return {string}
 */
func (str String) Trim(set string) string {
	return strings.Trim(string(str), set)
}

// Retorna el TypeCode del dato
func (str String) TypeCode() TypeCode {
	return _STRING
}

// interface IComparer
func (str String) Compare(z string) int {
	return strings.Compare(string(str), z)
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (str String) CompareTo(vInt32 string) string {
	if w, ok := T(str).(IComparerS); ok {
		return w.Compare(vInt32)
	}
	return Literals().NullString
}

// interface IComparerIf
func (str String) CompareIf(vInt32 T) int {
	return strings.Compare(string(str), vInt32.(string))
}

// Compara dos valores. Da 1, 0, -1 según sea >, ==, < que el otro valor.
func (str String) CompareToIf(vInt32 T) string {
	if w, ok := T(str).(IComparerIfS); ok {
		return w.Compare(vInt32)
	}
	return Literals().NullString
}

// interface IEquater
func (str String) Equals(vInt32 string) bool {
	return string(str) == vInt32
}

// Compara dos valores. Da true, false.
func (str String) EqualsTo(vInt32 string) bool {
	if w, ok := T(str).(IEquaterS); ok {
		return w.Equals(vInt32)
	}
	return false
}

// interface IEquaterIf
func (str String) EqualsIf(vInt32 T) bool {
	return string(str) == vInt32.(string)
}

// Compara dos valores. Da true, false.
func (str String) EqualsToIf(vInt32 T) bool {
	if w, ok := T(str).(IEquaterIf); ok {
		return w.Equals(vInt32)
	}
	return false
}
