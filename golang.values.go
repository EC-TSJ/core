/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

// type Values map[string]UNumber

// var (
// 	// GlobalValues
// 	GlobalValues Values = Values{
// 		// "EC-TSJ_PATH": 0,
// 	}

// 	//!+
// 	__values Values = make(Values)
// )

// // Ancho de GlobalValues
// // ! @return {int}
// func (v Values) Length() int {
// 	return len(GlobalValues)
// }

// // Retorna un valor de GlobalValues
// // ! @return {int}
// func (v Values) Get(s string) UNumber {
// 	return GlobalValues[s]
// }

// /**
//   - Crea un registro único, suma de los registros globales, locales, etc
//   - Los valores de los registros de la derecha tienen preferencia sobre los valores de la izquierda, se
//   - han de poner primero globales y luego locales.
//   - <CODE>
//     *?		golang.MakeValues(GlobalValues, LocalValues)
//   - </CODE>
//     *! @param{...map[string]int} los registros a sumar
//     */
// func MakeValues(m /*...Values*/ Values) {
// 	__values := maps.Clone(GlobalValues)
// 	maps.Copy(__values, m)
// 	// var (
// 	// 	n int = 0
// 	// )
// 	// m = append([]Values{GlobalValues}, m[0:]...)
// 	// for n < len(m) {
// 	// 	for k, v := range m[n] {
// 	// 		__values[k] = v
// 	// 	}
// 	// 	n++
// 	// }
// }

// /**
//  * Obtiene un valor del registro de valores
//  *! @param{string}
//  *! @return {any}
//  */
// func GetValue(s string) UNumber {
// 	return __values[s]
// }

// //!-

// /**
//  * Obtiene el GlobalValues
//  *! @return {Registry}
//  */
// func GValues() Values {
// 	return GlobalValues
// }
