/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"unsafe"
)

var ()

const ()

type (
	// Datos genéricos
	T interface{}
	// Dato de valor
	Value T
	// Dato de clave
	Key T
	// Tipo literals
	literals struct {
		NullString     string
		NewLine        string
		LF             string
		CarriageReturn string
		CR             string
		Alert          string
		BackSpace      string
		FormFeed       string
		FF             string
		HorizontalTab  string
		VerticalTab    string
		Backslash      string
		SingleQuote    string
		DoubleQuote    string
		Space          string
	}

	// Tipo Parm
	Parm struct {
		Void unsafe.Pointer
		What reflect.Kind // "byte", "int8", "uint8","int16", "uint16", "int32", "float32", "rune", "uint32"
		// "int", "float64", "int64", "complex64", "uint", "uint64", "string", "complex128"
	}

	// Tipo Pair
	Pair struct {
		T1 T
		T2 T
	}

	// Código Base de Error y Warning
	BaseError struct {
		Code int
		Err  string
		Text string
	}

	// Warning. Hereda de BaseError
	Warning struct {
		BaseError
	}

	// bloque de try...catch
	TryCatch struct {
		Try     func()
		Catch   func(Exception)
		Finally func()
	}

	// exception for catch
	Exception T

	// bloque del While..., DoWhile... y RepeatUntil...
	While struct {
		Condition func() bool
		Body      func()
	}

	// bloque del if
	If struct {
		Condition func() bool
		TrueOP    func() T
		FalseOP   func() T
	}

	// punteros
	fnPairCallback func(T, T) *Pair

	// // decorador
	// Decorate struct {
	// 	Decorator func(func(...T) T, ...T) T
	// 	Decorable func(...T) T
	// }

	Decorate struct {
		Decorator func(...T) func(...T) T
		Decorable func(...T) T
	}
)

/**
 * Decorador
 */
// func (this Decorate) Decorate(args ...T) T {
// 	return this.Decorator(this.Decorable, args...)
// }

/**
 * Decorador
 */
func (this Decorate) Decorate(args ...T) T {
	return (this.Decorator(this.Decorable(args...)))() // en una línea
}

//!+
/**
  * Nos da los literales de runas existentes
	* @return {*literals}
	* @@obsolete
*/
func Literals() *literals {
	return &literals{NullString: "", NewLine: "\n", LF: "\n", CarriageReturn: "\r", CR: "\r", Alert: "\a",
		BackSpace: "\b", FormFeed: "\f", FF: "\f", HorizontalTab: "\t", VerticalTab: "\v", Backslash: "\\",
		SingleQuote: "'", DoubleQuote: "\"", Space: " "}
}

/**
 * Nos da los literales de runas existentes
 */
var Lit *literals = &literals{NullString: "", NewLine: "\n", LF: "\n", CarriageReturn: "\r", CR: "\r", Alert: "\a",
	BackSpace: "\b", FormFeed: "\f", FF: "\f", HorizontalTab: "\t", VerticalTab: "\v", Backslash: "\\",
	SingleQuote: "'", DoubleQuote: "\"", Space: " "}

var L *literals = &literals{NullString: "", NewLine: "\n", LF: "\n", CarriageReturn: "\r", CR: "\r", Alert: "\a",
	BackSpace: "\b", FormFeed: "\f", FF: "\f", HorizontalTab: "\t", VerticalTab: "\v", Backslash: "\\",
	SingleQuote: "'", DoubleQuote: "\"", Space: " "}

//!-

// Interface error
func (e *BaseError) Error() string {
	return fmt.Sprintf("BaseError: [%d]'%s', '%s'.", e.Code, e.Err, e.Text)
}

// Interface error
func (w *Warning) Error() string {
	return fmt.Sprintf("Warning: [%d]'%s', '%s'.", w.Code, w.Err, w.Text)
}

// NewPair: Obtiene un nuevo Par
// @param {T}
// @param {T}
// @return {Pair}
func (*Pair) New(t1, t2 T) *Pair {
	return &Pair{T1: t1, T2: t2}
}

// NewPair: Obtiene un nuevo Par
// @@obsolete
var NewPair fnPairCallback = (&Pair{}).New

// MakePair: Obtiene un nuevo Par
// @param {T}
// @param {T}
// @return {Pair}
var MakePair fnPairCallback = (&Pair{}).New

//!+
/**
 * Evalúa el operador ternario de C, el ?:.
 * Debe ser usado con type assertion
 *  <CODE>
 *     ...
 *	 	   parm := core.IIf(uintptr, uintptr, 0).(uintptr)
 *     ...
 *
 *  </CODE>
 * @param {bool} expresión a evaluar
 * @param {interface{}} operación True
 * @param {interface{}} operación False
 * @return {interface{}}
 */
func IIf(f bool, trueOp T, falseOp T) T {
	if f {
		return trueOp
	} else {
		return falseOp
	}
}

/**
 * Evalúa el operador ternario de C, el ?:.
 * <CODE>
 *   If {
 *		Condition: func() bool { return T == 25 }
 * 		TrueOP: func() interface{}{ return "clock" }
 * 		FalseOP: func() interface{}{ return "Flip" }
 *   }.If().(string)
 *</CODE>
 * @return {interface{}}
 */
func (this If) If() T {
	if this.Condition() {
		return this.TrueOP()
	} else {
		return this.FalseOP()
	}
}

/**
 * Evalúa el while{...} de C.
 * <CODE>
 *	While{
 *		Condition: func() bool { return T == 25 },
 *		Body:      func() { T++ },
 *	}.While()
 * </CODE>
 * the same that:
 * <ALTERNATIVE>
 *  for T == 25 {
 *    ...
 *		...
 * }
 * </ALTERNATIVE>
 */
func (this While) While() {
AnotherTime:
	if this.Condition() { // si la condición es true
		this.Body()
		goto AnotherTime
	}
}

/**
 * Evalúa el do{...}while de C.
 * <CODE>
 *	While{
 *		Condition: func() bool { return T == 25 },
 *		Body:      func() { T++ },
 *	}.DoWhile()
 * </CODE>
 * the same that:
 * <ALTERNATIVE 1>
 * 	for {
 *    ...
 *		...
 * 	  if !(T == 25) {
 *	    continue
 *    } else {
 *      break
 *    }
 * 	}
 * </ALTERNATIVE 1>
 * <ALTERNATIVE 2>
 *     enlace:  https://yourbasic.org/golang/do-while-loop/
 *   for ok := true; ok; ok = condition {
 *     work()
 *   }
 * </ALTERNATIVE 2>
 */
func (this While) DoWhile() {
AnotherTime:
	this.Body()
	if this.Condition() {
		goto AnotherTime // mientras que sea true
		// y hasta que sea falso
	}
}

/**
 * Evalúa el Repeat...Until... de Java.
 * <CODE>
 *	While{
 *		Condition: func() bool { return T == 25 },
 *		Body:      func() { T++ },
 *	}.RepeatUntil()
 * </CODE>
 * the same that:
 * <ALTERNATIVE 1>
 * 	for {
 *    ...
 *		...
 * 	  if T == 25 {
 *	    break
 *    }
 * 	}
 * </ALTERNATIVE 1>
 * <ALTERNATIVE 2>
 *     enlace: https://yourbasic.org/golang/do-while-loop/
 *   for ok := true; ok; ok = !condition {
 *     work()
 *   }
 * </ALTERNATIVE 2>
 */
func (this While) RepeatUntil() {
AnotherTime:
	this.Body()
	if this.Condition() {
		return // mientras que sea falso
		// hasta que sea true
	}
	goto AnotherTime
}

/**
 * Throw para el sistema try...catch
 * @param {Exception}
 */
func Throw(up Exception) {
	panic(up)
}

/**
 * Do. Para el sistema try...catch
 * <CODE>
 *   ...
 *    TryCatch{
 * 			Try: func() {
 *				Throw(NewError("JoderLaPana", 0x456, "ddddd ffffffff ggggg"))
 *			},
 *			Catch: func(e Exception) {
 *			  if fmt.Sprintf("%T", e) == "*BaseError" {
 *				  fmt.Println("yes")
 *			  }
 *				if (e.(*BaseError)).Err == "JoderLaPana" {
 *					fmt.Printf("%v\n", e)
 *				}
 *			},
 *			Finally: func() {
 *	 			fmt.Printf("Lo que sea.\n", )
 *      },
 *    }.TryCatch()
 *    ...
 * </CODE>
 */
func (this TryCatch) TryCatch() {
	if this.Finally != nil {
		defer this.Finally()
	}
	if this.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				this.Catch(r)
			}
		}()
	}
	this.Try()
}

/**
 * Evalúa el argumento opcional y nos lo devuelve, ó 0, ó false, ó cadena vacía, si es el caso.
 * El parámetro debe ser `...core.T` en la función que lo recibe (y que luego llama a ArgOptional).
 * Se deberá realizar una `type assertion` con .(string) ó .(int) ó lo que sea, posteriormente. Ejemplo:
 *  <CODE>
 *     func calling(msg ...interface{}) {
 *	 	   parm := ArgOptional("default", msg).(string)
 *       ...
 *     }
 *  </CODE>
 *  <ANOTHER POSIBILITY>
 *     func calling(s ...interface{}) {
 * 	     s = append(s, nil, nil, nil) // inicializa los parámetros
 *	     s[0] = map[bool]interface{}{true:  s[0], false: "cojones"}[s[0] != nil]
 *	     s[1] = map[bool]interface{}{true:  s[1], false: 6.31}[s[1] != nil]
 *	     s[2] = map[bool]interface{}{true:  s[2], false: 5}[s[2] != nil]
 *	     s = s[:3] // recorta el número de parámetros
 *       ...
 *     }
 *  </ANOTHER POSIBILITY>
 * @param {interface{}} valor default, debe ser del mismo tipo que el siguiente
 * @param {...interface{}} valor opcional, si existe, del mismo tipo que el anterior
 * @return {interface{}}
 */
func ArgOptional(_default T, _optional ...T) T {
	if _optional[0].([]T) == nil {
		return _default
	} else {
		return (_optional[0].([]T))[0]
	}
}

//!-
/**
 * Convierte una cadena de bytes a una string
 * @param {[]byte}
 * @return {string}
 */
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

/**
 * Modifica una string
 * <CODE>
 *     ...
 *     ModifyString(*cadena, 5, 3)
 *     ...
 * </CODE>
 * @param {*string} cadena a modificar
 * @param {uintptr} comienzo de la cadena
 * @param {int} final de la cadena
 */
func ModifyString(str *string, starting uintptr, ending int) {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(str))
	stringHeader.Data = stringHeader.Data + /*comienzo*/ IIf(starting > 0, starting, uintptr(0)).(uintptr)
	stringHeader.Len = stringHeader.Len - /*offset ó comienzo*/ IIf((int(starting)+ending) != 0, int(starting)+ /*final*/ ending, 0).(int)
}

/**
 * Modifica un slice
 * <CODE>
 *    ...
 *    ModiFySlice(&Parm{Void: Pointer(&T), What: reflect.String}, 5, 2)
 *    ...
 * </CODE>
 * @param {*Parm} slice a modificar
 * @param {uintptr} comienzo del slice
 * @param {int} final del slice
 */
func ModifySlice(parm *Parm, starting uintptr, ending int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(parm.Void))
	switch parm.What {
	case /* byte */ reflect.Int8, reflect.Uint8:
		ss.Data = ss.Data + /*comienzo*/ 1*IIf(starting > 0, starting, uintptr(0)).(uintptr)
	case reflect.Int16, reflect.Uint16:
		ss.Data = ss.Data + /*comienzo*/ 2*IIf(starting > 0, starting, uintptr(0)).(uintptr)
	case /* rune */ reflect.Int32, reflect.Float32, reflect.Uint32:
		ss.Data = ss.Data + /*comienzo*/ 4*IIf(starting > 0, starting, uintptr(0)).(uintptr)
	case reflect.Int, reflect.Float64, reflect.Int64, reflect.Complex64, reflect.Uint, reflect.Uint64:
		ss.Data = ss.Data + /*comienzo*/ 8*IIf(starting > 0, starting, uintptr(0)).(uintptr)
	case reflect.String, reflect.Complex128:
		ss.Data = ss.Data + /*comienzo*/ 16*IIf(starting > 0, starting, uintptr(0)).(uintptr)
	}
	ss.Len = ss.Len - /*offset ó comienzo*/ IIf((int(starting)+ending) != 0, int(starting)+ /*final*/ ending, 0).(int)
	ss.Cap = ss.Len + 5
}

//!+

type Uuid [16]byte

// Uuid: Obtiene un Uuid
// @return {_uuid_}
func UUID() (u *Uuid) {
	u = new(Uuid)
	rand.Read(u[:])
	u[8] = (u[8] | 0x40) & 0x7F    // setVariant - 0x40
	u[6] = (u[6] & 0xF) | (4 << 4) // setVersion - 4
	return
}

// Retorna version desparseada de la secuencia Uuid.
// interface Stringer
func (u *Uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//!-
