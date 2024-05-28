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
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

type (
	//?  @public @union
	Union struct {
		__base__ unsafe.Pointer
	}

	// Tipo Parm
	//? @class
	Parm struct {
		Void unsafe.Pointer
		What reflect.Kind // "byte", "int8", "uint8","int16", "uint16", "int32", "float32", "rune", "uint32"
		// "int", "float64", "int64", "complex64", "uint", "uint64", "string", "complex128"
	}

	// para la sobrecarga de metodos
	//? @class
	Ovl1 struct {
		BytesToString FuncV1[[]byte, string] //func(b []byte) string
		StringToBytes FuncV1[string, []byte]
		ModifyString  ActionV3[*string, uintptr, int] //func(str *string, starting uintptr, ending int)
		ModifySlice   ActionV3[*Parm, uintptr, int]   //func(parm *Parm, starting uintptr, ending int)
		Statically    FuncV1[int, func() int]
	}

	//? @class
	Ovl2[T Ordered] struct {
		MaxOf       func([]T) T
		MinOf       func([]T) T
		ArgOptional FuncV2X[T, any, T]
	}

	// For use with MethodWithModifiers and other types of functions
	//? @class
	Return struct {
		MethodType MethodType
		Return     any
		Error      error
	}
)

//
// The modified method is some like this:
//
//? func MethodWithModifiers(methodType MethodType) func(...any) *Return {
// 	return func(a ...any) *Return {
// 		a[0] = a[0].(int) + int(methodType)
// 		if reflect.TypeOf(a[0]).Kind() == reflect.Int {
// 			return &Return{methodType, a[0].(int), nil}
// 		}
// 		return &Return{methodType, -1, nil}
// 	}
// }

func (*Ovl1) New() *Ovl1 {
	return &Ovl1{
		Statically:
		/**
		 * Nos da un valor de tipo Static. Del tipo static de C, no de las clases de instancia/valor
		 * de Java, p. ej.
		 *
		 *? <CODE>
		 *?		f := core.Statically(0) // f() => 0
		 *?		fmt.Println(  f()  ) //0
		 *?		fmt.Println(  f()  ) //1
		 *?		fmt.Println(  f()  ) //2
		 *?		f = core.Statically(51) // f() => 51
		 *?		fmt.Println(  f()  ) //51
		 *?		fmt.Println(  f()  ) //52
		 *?		fmt.Println(  f()  ) //53
		 *? </CODE>
		 *
		 *! @param {...T} valor base
		 *! @return {func() T}
		 */
		func(val int) func() int {
			return func() int {
				p := val
				val++
				return p
			}
		},

		BytesToString:
		//!-
		/**
		 * Convierte una cadena de bytes a una string
		 *
		 * !@param {[]byte}
		 * !@return {string}
		 */
		func(b []byte) string {
			if len(b) == 0 {
				return L.NullString
			}
			return unsafe.String(unsafe.SliceData(b), len(b))
		},

		StringToBytes: func(s string) []byte {
			if s == L.NullString {
				return nil
			}
			return unsafe.Slice(unsafe.StringData(s), len(s))
		},

		ModifyString:
		/**
		* Modifica una string
		*
		* <CODE>
		*?     ...
		*?     ModifyString(*cadena, 5, 3)
		*?     ...
		* </CODE>
		*
		* !@param {*string} cadena a modificar
		* !@param {uintptr} comienzo de la cadena
		* !@param {int} final de la cadena
		 */
		func(str *string, starting uintptr, ending int) {
			stringHeader := (*reflect.StringHeader)(unsafe.Pointer(str))
			stringHeader.Data = stringHeader.Data + /*comienzo*/ If[uintptr](starting > 0)(starting, uintptr(0))
			stringHeader.Len = stringHeader.Len - /*offset ó comienzo*/ If[int]((int(starting)+ending) != 0)(int(starting)+ /*final*/ ending, 0)
		},

		ModifySlice:
		/**
		 * Modifica un slice
		 *
		 * <CODE>
		 *?    ...
		 *?    ModiFySlice(&Parm{Void: Pointer(&T), What: reflect.String}, 5, 2)
		 *?    ...
		 * </CODE>
		 *
		 * !@param {*Parm} slice a modificar
		 * !@param {uintptr} comienzo del slice
		 * !@param {int} final del slice
		 */
		func(parm *Parm, starting uintptr, ending int) {
			ss := (*reflect.SliceHeader)(unsafe.Pointer(parm.Void))
			switch parm.What {
			case /* byte */ reflect.Int8, reflect.Uint8:
				ss.Data = ss.Data + /*comienzo*/ 1*If[uintptr](starting > 0)(starting, uintptr(0))
			case reflect.Int16, reflect.Uint16:
				ss.Data = ss.Data + /*comienzo*/ 2*If[uintptr](starting > 0)(starting, uintptr(0))
			case /* rune */ reflect.Int32, reflect.Float32, reflect.Uint32:
				ss.Data = ss.Data + /*comienzo*/ 4*If[uintptr](starting > 0)(starting, uintptr(0))
			case reflect.Int, reflect.Float64, reflect.Int64, reflect.Complex64, reflect.Uint, reflect.Uint64:
				ss.Data = ss.Data + /*comienzo*/ 8*If[uintptr](starting > 0)(starting, uintptr(0))
			case reflect.String, reflect.Complex128:
				ss.Data = ss.Data + /*comienzo*/ 16*If[uintptr](starting > 0)(starting, uintptr(0))
			}
			ss.Len = ss.Len - /*offset ó comienzo*/ If[int]((int(starting)+ending) != 0)(int(starting)+ /*final*/ ending, 0)
			ss.Cap = ss.Len + 5
		},
	}
}

// ! @overloaded methods
func (*Ovl2[T]) New() *Ovl2[T] {
	return &Ovl2[T]{
		MaxOf:
		/*************************************************************************
		 * Nos dice el número que es mayor
		 *
		 *! param{T}
		 *! param {T}
		 *! return {T}
		 */
		func(t []T) T {
			f := t[0]
			for _, v := range t {
				if f < v {
					f = v
				}
			}
			return f
		},

		MinOf:
		/*************************************************************************
		 * Nos dice el número que es mayor
		 *
		 *! param{T}
		 *! param {T}
		 *! return {T}
		 */
		func(t []T) T {
			f := t[0]
			for _, v := range t {
				if f > v {
					f = v
				}
			}
			return f
		},

		ArgOptional:
		/***************************************************************************
		 * Evalúa el argumento opcional y nos lo devuelve, ó 0, ó false, ó cadena vacía, si es el caso.
		 * El parámetro debe ser `...core.T` en la función que lo recibe (y que luego llama a ArgOptional).
		 *  Ha de llamarse internamente desde una función, que tenga un parámetro ...any.
		 *
		 * Ejemplo:
		 *  <CODE>
		 *?     func calling(msg ...any) {
		 *?	 	   parm := (&Ovl2[string]{}).ArgOptional("default", msg)
		 *?      ...
		 *     }
		 *  </CODE>
		 *  <ANOTHER POSIBILITY>
		 *?     func calling(s ...interface{}) {
		 *? 	     s = append(s, nil, nil, nil) // inicializa los parámetros
		 *?	     s[0] = map[bool]interface{}{true:  s[0], false: "cojones"}[s[0] != nil]
		 *?	     s[1] = map[bool]interface{}{true:  s[1], false: 6.31}[s[1] != nil]
		 *?	     s[2] = map[bool]interface{}{true:  s[2], false: 5}[s[2] != nil]
		 *?	     s = s[:3] // recorta el número de parámetros
		 *?       ...
		 *?     }
		 *  </ANOTHER POSIBILITY>
		 * !@param {T} valor default, debe ser del mismo tipo que el siguiente
		 * !@param {...interface{}} valor opcional, si existe, del mismo tipo que el anterior
		 * !@return {T}
		 **************************************************************************************************
		 *! @ArgOptional(T, ...T) T
		 */
		func(_default T, _optional ...any) T {
			if len(_optional[0].([]any)) > 0 {
				return _optional[0].([]any)[0].(T)
			}
			return _default
		},
	}
}

// helper function
func _validate(fnType reflect.Type, tf bool) {
	if fnType.Kind() != reflect.Func { // es funcion
		panic("Error: debe ser una función")
	}
	if tf {
		if fnType.NumIn() != 0 { // si nº parametros != 1 ó  el tipo del parámetro 0 implementa Error
			panic("Error: no debe tener ningún argumento")
		}
	} else {
		if fnType.NumIn() != 1 { // si nº parametros != 1 ó  el tipo del parámetro 0 implementa Error
			panic("Error: debe tener sólo un argumento")
		}
	}
	if fnType.NumOut() != 0 { // si nº de retornos no es 0
		panic("Error: Resultado: no debe tener ningún resultado")
	}
}

//!+ Lit
/**
 * Nos da una enum de los literales de runas existentes
 *
 *? @public @enum
 */
var Lit = &struct {
	NullString, NewLine, LF, CarriageReturn, Colon, SemiColon, CR, Alert, BackSpace, FormFeed, FF, HorizontalTab,
	VerticalTab, Backslash, SingleQuote, DoubleQuote, Space string
}{
	NullString:     "",
	NewLine:        "\n",
	LF:             "\n",
	CarriageReturn: "\r",
	Colon:          ":",
	SemiColon:      ";",
	CR:             "\r",
	Alert:          "\a",
	BackSpace:      "\b",
	FormFeed:       "\f",
	FF:             "\f",
	HorizontalTab:  "\t",
	VerticalTab:    "\v",
	Backslash:      "\\",
	SingleQuote:    "'",
	DoubleQuote:    "\"",
	Space:          " ",
}

//!-

var (
	L             = Lit
	BytesToString = (&Ovl1{}).New().BytesToString //func(b []byte) string
	StringToBytes = (&Ovl1{}).New().StringToBytes
	ModifyString  = (&Ovl1{}).New().ModifyString //func(str *string, starting uintptr, ending int)
	ModifySlice   = (&Ovl1{}).New().ModifySlice  //func(parm *Parm, starting uintptr, ending int)
	Statically    = (&Ovl1{}).New().Statically
	/** Genericos */
	ArgOptionalS = (&Ovl2[string]{}).New().ArgOptional  //func(T, ...T)  T
	ArgOptionalF = (&Ovl2[float64]{}).New().ArgOptional //func(T, ...T)  T
	ArgOptionalI = (&Ovl2[int]{}).New().ArgOptional     //func(T, ...T)  T
	MinOf        = (&Ovl2[int]{}).New().MinOf           //func([]T)  T
	MaxOf        = (&Ovl2[int]{}).New().MaxOf           //func([]T)  T
)

/**
 * Throw para el sistema try...catch
 *
 * !@param {Exception}
 */
func Throw(err error) bool {
	panic(err)
}

/**
 * Sistema try...catch
 *! @param {func()}
 *! @param {func()}
 */
func Exception(_recover func()) func(_panic func()) {
	return func(_panic func()) {
		if _recover != nil {
			defer _recover()
		}
		_panic()
	}
}

/**
 * TryCatch. Para el sistema try...catch
 *
 * <CODE>
 *   ...
 *?    Try(func() {
 *?			   // codigo regular que produce excepciones
 *?				Throw(BaseError("JoderLaPana", 0x456, "ddddd ffffffff ggggg"))
 *?				// más código regular
 *?			}) ([]any{		func(e *errors.BaseError) {
 *?				if errors.Is(e, errors.Err(errors.GetError(errors.Name(e)), "Correct")) {
 *?					fmt.Println("BE> ", e)
 *?				}
 *?			}, func(e *errors.Warning) {
 *?				if errors.Is(e, errors.Err(errors.GetError(errors.Name(e)), "Warning")) {
 *?					fmt.Println("W> ", e)
 *?				}
 *?			}, func(e *errors.HTTPError) {
 *?				if errors.Is(e, errors.Err(errors.GetError(errors.Name(e)), "Warning")) {
 *?					fmt.Println("HE> ", e)
 *?				}
 *?			}, func(e error) {
 *?				fmt.Println("PE> SOY PATHERROR", e)
 *?			},
 *?		 }) (func() {
 *?	  			fmt.Printf("Lo que sea.\n", )
 *?       })
 *    ...
 * </CODE>
 */
func Try(try func()) func(catch []any) func(finally ...func()) {
	return func(catch []any) func(finally ...func()) {
		return func(finally ...func()) {
			if finally != nil {
				defer finally[0]()
			}
			if catch != nil {
				defer func() {
					count, idx := len(catch), 0
					if rec := recover(); rec != nil {
						for idx < count {
							fnType := reflect.TypeOf(catch[idx])
							if fnType.NumIn() == 1 {
								if reflect.TypeOf(rec.(error)).AssignableTo(fnType.In(0)) {
									reflect.ValueOf(catch[idx]).Call([]reflect.Value{reflect.ValueOf(rec.(error))})
								}
							}
							idx++
						}
					}
				}()
			}
			try()
		}
	}
}

/**
 * Evalúa el while{...} de C.
 *
 * <CODE>
 *?   While(T == 26)(func(...any) bool {
 *?	    ...
 *?     return true ó false
 *?   })
 * </CODE>
 * the same that:
 * <ALTERNATIVE>
 *?  for T == 25 {
 *?    ...
 *?		 ...
 *? }
 * </ALTERNATIVE>
 */
func While(flag bool) func(func(...any) bool, ...any) {
	return func(fn func(...any) bool, parameters ...any) {
		for flag {
			flag = fn(parameters...)
		}
	}
}

/**
 * Evalúa el do{...}while de C.
 *
 * <CODE>
 *?   DoWhile(T == 26)(func(...any) bool {
 *?	    ...
 *?     return true ó false
 *?   })
 * </CODE>
 * the same that:
 * <ALTERNATIVE 1>
 *? 	for {
 *?    ...
 *?		 ...
 *?	   if !(T == 25) {
 *?	     continue
 *?    } else {
 *?      break
 *?    }
 *? 	}
 * </ALTERNATIVE 1>
 * <ALTERNATIVE 2>
 *     enlace:  https://yourbasic.org/golang/do-while-loop/
 *?   for ok := true; ok; ok = condition {
 *?     work()
 *?   }
 * </ALTERNATIVE 2>
 */
func DoWhile(flag bool) func(body func(args ...any) bool, parameters ...any) {
	return func(fn func(args ...any) bool, parameters ...any) {
		flag = fn(parameters...)
		for flag {
			flag = fn(parameters...)
		}
	}
}

//!+ Interpolation
/**
 * Por defecto para las interpolaciones
 */
var DefaultMapOfInterpolations map[string]string
var W FuncV3X[string, int, map[string]string, string] = Interpolation
var __fileName string

/**
 * Realiza la interpolación de cadenas. Necesita un DefaultMapOfInterpolations, en su caso, para que funcione.
 * Debe ser usado desde la función fmt.Printf(), del modulo std/fmt (no el modulo fmt).
 * Orden de interpolación:
 *  - A) Lista de todas las variables y constantes de tipo token.STRING (vars & consts string)
 *  - B) Variables de Entorno
 *  - C) Variable por defecto (la expansion de ${ENV:var} )
 *  - D) DefaultMapOfInterpolations variable (or _map value). NO, for PRINTF calls
 *
 *! @param {string}
 *! @param {int}
 *! @param {map[string]string}
 *! @return {string}
 */
func Interpolation(s string, nf /** 1 or 2 */ int, body ...map[string]string) string {
	if __fileName == L.NullString {
		_, __fileName, _, _ = runtime.Caller(nf /* is the file backward two jumps (callings to functions) */)
	}
	var (
		char, ends, corp string
		led              bool
	)
	const (
		LeftBracket  string = "${"
		RightBracket string = "}"
	)
	// assignment
	if body == nil {
		body = make([]map[string]string, 1)
		body[0] = DefaultMapOfInterpolations
	}
	//
	count := len(strings.Split(s, LeftBracket)) - 1
	start := strings.SplitN(s, LeftBracket, 2) // first code
	if len(start) == 1 {                       // return to base
		return s
	} else { // second code
		end := strings.SplitN(start[1], RightBracket, 2)
		instr := strings.Split(end[0], L.Colon) // for search default value inline
		if len(instr) == 1 {
			char, led = strings.Trim(instr[0], L.Space), false // is tag value
		} else {
			char, led = strings.Trim(instr[1], L.Space), true // is default value for set
		}
		ends = end[1] // second code. final part
	}
	/** Interpolate's Helper function. Format string */
	var __sep = func(first, core, last string) string {
		return first + core + last
	}
	/** Interpolate's Helper function. *Ast File */
	var __panicking = func(filename string) *ast.File {
		fast, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		return fast
	}
	/** Interpolate's Helper Function. Search vars for a string of equal name */
	var __varSearch = func(s, filename string) string {
		fast := __panicking(filename)
		for _, d := range fast.Decls { // busca en las declaraciones
			switch decl := d.(type) {
			case *ast.GenDecl:
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {
					case *ast.ValueSpec:
						for _, id := range spec.Names {
							if id.Name == s { // es la variable buscada
								lt := id.Obj.Decl.(*ast.ValueSpec).Values[0]
								if v, ok := lt.(*ast.BasicLit); ok { // Valores primitivos
									if v.Kind == token.STRING { // Valores que son STRING
										return v.Value[1 : len(v.Value)-1] // quita las comillas dobles
									}
								}
							}
						}
					}
				}
			}
		}
		return L.NullString
	}
	// see vars for a string of equal name
	corp = __varSearch(char, __fileName)
	if corp == L.NullString { // Env or default variables
		corp = os.Getenv(char)
		if corp == L.NullString {
			corp = If[string](!led)(body[0][char] /* led = false; map value */, char /* led = true; default value */)
		}
	}
	// epiface. Loop for get others interpolations
	var str = __sep(start[0], corp, ends)
	count--
	/* While (count > 0) */
	for count > 0 {
		str = Interpolation(str, nf, body...)
		count--
	}
	// save to nullstring
	__fileName = L.NullString
	return str //save str
}

// !-
/**
 * Nos dice si un tipo 'instance' contiene otro de tipo 'root'. Es decir, si 'root' esta contenido en 'instance'.
 * O, dicho de otro modo, si 'instance' es instancia de 'root'.
 *
 * Uso:
 * <CODE>
 *? type S struct {
 *?	  zS int
 *? }
 *?
 *? type M[T Ordered] struct {
 *?		*S
 *?		wM string
 *?	}
 *?
 *? type Example struct {
 *?		*D
 *?		A int
 *?		B struct {
 *?		  C int
 *?		  D **struct {
 *?			  E string
 *?			}
 *?		}
 *?		*M[int]
 *?	}
 *?
 *? flag := InstanceOf(reflect.TypeOf(&Example{}), reflect.TypeOf(&S{}))
 *? flag := InstanceOf(reflect.TypeOf(&Example{}), reflect.TypeOf((int)(0)))
 *? flag := InstanceOf(reflect.TypeOf(&Example{}), reflect.TypeOf((*int)(nil)))
 * </CODE>
 *
 *! @param {reflect.Type} instance. el contenedor (la instancia)
 *! @param {reflect.Type} root. el que está contenido ó raiz
 *! @return {bool}
 */
func InstanceOf(instance, root reflect.Type) bool {
	if instance == root { // controla los tipos tal cual
		return true
	}
	//
	switch instance.Kind() {
	case reflect.Pointer: // mira si es un puntero, entonces pasa a la instancia sin ser puntero
		flag := InstanceOf(instance.Elem(), root)
		if flag {
			return flag
		}
	case reflect.Struct: // esto es la instancia, no ptr
		for i := 0; i < instance.NumField(); i++ {
			if instance.Field(i).Type == root {
				return true
			}
			flag := InstanceOf(instance.Field(i).Type, root)
			if flag {
				return flag
			}
		}
	}
	return false
}

// !+ UUID
// ? @public @class Uuid
type Uuid [16]byte

// Uuid: Obtiene un Uuid
// !@return {_uuid_}
func UUID() (u *Uuid) {
	u = new(Uuid)
	rand.Read(u[:])
	u[8] = (u[8] | 0x40) & 0x7F    // setVariant - 0x40
	u[6] = (u[6] & 0xF) | (4 << 4) // setVersion - 4
	return
}

// Retorna version desparseada de la secuencia Uuid.
// !interface Stringer
func (u *Uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//!-
