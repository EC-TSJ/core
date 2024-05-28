/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"fmt"
	"reflect"
	"strings"
)

//!+

/**
 * Obtiene una lista de las propiedades de un tipo
 *
 * !@param {any} a POINTER data
 * !@return {[]string}
 * Luego se puede hacer ([]reflect.Value).Kind(),
 * ([]reflect.Value).Interface(), ([]reflect.Value).Type(),
 * ([]reflect.Value).etc
 */
func GetProperties(element any) ([]string, []reflect.Value) {
	tt := make([]string, 0)
	nn := make([]reflect.Value, 0)

	v := reflect.ValueOf(element).Elem()
	for j := 0; j < v.NumField(); j++ {
		name := v.Type().Field(j).Name
		rf := v.Field(j)
		tt = append(tt, name)
		nn = append(nn, rf)
	}

	return tt, nn
}

/**
 * Obtiene una lista de los métodos de un tipo
 *
 * !@param {any} a POINTER to data
 * !@return {[]string}
 */
func GetMethods(element any) []string {
	t := make([]string, 0)
	reflected := reflect.TypeOf(element)
	for element := 0; element < reflected.NumMethod(); element++ {
		m := reflected.Method(element)
		t = append(t, m.Name)
	}

	return t
}

/**
 * Especifica el valor como un any
 *
 * !@param {any}
 * !@return {any}
 */
func Interface(element any) any {
	return reflect.ValueOf(element).Interface()
}

/**
 * Especifica el tipo (reflect.Kind) de un elemento en forma de enumeración Kind
 * !@param {any}
 * !@return {reflect.Kind}
 */
func Kind(element any) reflect.Kind {
	return reflect.ValueOf(element).Kind()
}

/**
 * TypeSprintf nos dice de que tipo es un elemento dado. Nos lo dice mediante
 * una cadena del tipo. *** TypeOf() ***
 *
 * !@param {any} element
 * !@return {string}
 *
func TypeSprintf(element any) string {
	return fmt.Sprintf("%T", element)
}
*/

/**
 * Nos dice el nombre del tipo
 *
 *! @param {any}
 *! @return {string}
 */
func TypeOf(val any) string {
	v := fmt.Sprintf("%T", val)
	v1 := strings.Split(v, "main.")
	if len(v1) == 1 {
		return v1[0]
	} else {
		return v1[0] + v1[1]
	}
}

/**
 *TypeOf nos dice de que tipo es un elemento dado. Nos lo dice mediante TypeOf de reflect.
 *
 * !@param {any} element
 * !@return string
 */
// func TypeOf(element any) string {
// 	return reflect.TypeOf(element).String()
// }

/**
 * TypeIFace nos dice de que tipo es un elemento dado. Nos lo dice mediante
 * un switch. *** TypeOf() ***
 *
 * !@param {any} element
 * !@return string
 *
func TypeIFace(element any) any {
	switch v := element.(type) {
	default:
		return v
	}
}
*/
/**
 * Obtiene el nombre de un tipo (con * si se trata de un puntero)
 *   La diferencia con TypeOf es que no da el paquete en el nombre del
 * tipo. P. Ej. *core.Int32 da "*Int32" y no "*core.Int32" como daría
 * TypeOf.
 *
 * !@param {any} a POINTER to data
 * !@return {string}
 */
func GetType(element any) string {
	reflected := reflect.TypeOf(element)

	if reflected.Kind() == reflect.Ptr {
		return "*" + reflected.Elem().Name()
	} else {
		return reflected.Name()
	}
}

/**
 * Llama a una función por el tipo y nombre (y los argumentos)
 *
 *? <CODE>
 *?    type TypeOne struct{}
 *?    func (t *TypeOne) FuncOne() { }
 *?    func (t *TypeOne) FuncTwo(name string) { }
 *?    ...
 *?    ...
 *?    t1 := &TypeOne{}
 *?    ...
 *?    out, err := CallFuncByName(t1, "FuncOne")
 *?    if err != nil {
 *? 	    panic(err)
 *?    }
 *?     *Return value
 *?    _ = out
 *?    ...
 *?    ...
 *?    out, err = CallFuncByName(t1, "FuncTwo", "parameter")
 *?    if err != nil {
 *? 	    panic(err)
 *?    }
 *?     *return value
 *?    _  = out
 *? </CODE>
 * !@param {any}
 * !@param {string}
 * !@param {...any}
 * !@return {[]reflect.Value}
 * !@return {error}
 */
func CallFuncByName( /* tipo */ myClass any /* nombre funcion */, funcName string /* argumentos de la función */, params ...any) (out []reflect.Value, err error) {
	myClassValue := reflect.ValueOf(myClass)
	m := myClassValue.MethodByName(funcName)
	if !m.IsValid() {
		return make([]reflect.Value, 0) /*(&errors.Errors{}).Custom("Methods not found", L.NullString, 0x2009, funcName)*/, fmt.Errorf("Method not found \"%s\"", funcName)
	}
	in := make([]reflect.Value, len(params))
	for element, param := range params {
		in[element] = reflect.ValueOf(param)
	}
	out = m.Call(in)
	return
}

//!-

// type ssdd struct {
// }

// func (s *ssdd) Get() {}
// func (s *ssdd) Pd()  {}

// func main() {
// 	pd := ssdd{}
// 	name := GetType(&pd)
// 	methods := GetMethods(&pd)
// 	properties := GetProperties(pd)
// 	fmt.Println(name, methods, properties)
// 	ad := TypeOf(pd)
// 	ad1 := Type(pd)
// 	ad2 := TypeFrom(pd).(ssdd)
// 	fmt.Println(ad, ad1, ad2)
// }
