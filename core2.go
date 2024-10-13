/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

var version = "2.000.00"

/**
 * TODO(NOTA):
 * La diferencia entre las funciones 'ArgOptional{S|F|I}' y 'Option' es que la primera se usa internamente para encontrar un argumento
 * opcional de la función y se pasa el argumento que queremos encontrar enteramente como v, no como v[0], ambos son de tipo ...any, el de la
 * función llamada, y el de la llamante.
 * la segunda se usa externamente a cualquier función y se le llama por el tipo primitivo. Así:
 * la primera)
 *?     func calling[string](msg ...T) {
 *?	 	    parm := ArgOptional[string]("default", msg[0])
 *?       ...
 * la segunda)
 *?       Option[string]("default", "valor")
 *?       ...
 *
 * Realmente 'ArgOptional{S|F|I}' es como hacer 'UnInterface' y 'Option' después. Así:
 *?     func calling[string](v ...T) {
 *? 	    vv := core.UnInterface[string](v)
 *?	      core.Option[string](def, vv) ó
 *?       - core.Option[string]("default", core.UnInterface[string](v))
 */

/**
 * El operador ternario en una sola instrucción
 *
 * <CODE>
 *?  ...
 *?  If[int](5 > 7)(5, 7)
 *?  ...
 * </CODE>
 * ! @type {T comparable}
 * ! @param {bool}
 * ! @return {func(T, T) T}
 * ! @param {T}
 * ! @param {T}
 * ! @return {}
 */
func If[T comparable](flag bool) func(True T, False T) T {
	f := flag
	return func(True, False T) T {
		if f {
			return True
		}
		return False
	}
}

/***************************************************************************
 */
/***************************************************************************
 * Es el operador ternario ?: ó la función if
 *
 * !@type {T any}
 * !@param {bool}
 * !@param {func() T}
 * !@param {func() T}
 * !@returns {T}
 ********************************************
 * ! @Iif(f bool, a func() T,  b func() T) T
 */
func IIf[T comparable](f bool) func(func(...T) T, func(...T) T, ...T) T {
	flag := f
	return func(t func(args ...T) T, f func(args ...T) T, parms ...T) T {
		if flag {
			return t(parms...)
		}
		return f(parms...)
	}
}

/**
 * Convierte un array de un función variádica a un elemento utilizable
 *
 * ! @type {T comparable}
 * ! @param {any} el elemento que se le pasa es el array, no un elemento del array
 * !							es decir, v no v[0]
 * ! @return {T}
 */
func V[T comparable](v any) T {
	if v.([]T) == nil {
		return *new(T)
	}
	return v.([]T)[0]
}

/**
 * Desinterfacea un interface, un nivel
 *
 * ! @type {T any}
 * ! @param {any}
 * ! @return (T)
 */
func UnInterface[T comparable](data any) T {
	if len(data.([]any)) > 0 {
		return data.([]any)[0].(T)
	}
	return *new(T)
}

/**
 * Evalúa el argumento opcional y nos lo devuelve, ó 0, ó false, ó cadena vacía, si es el caso.
 *   Esta función es para ser llamada desde fuera de una función.
 * Para ser llamada desde una función se usa 'ArgOptional{S|I|F}'.
 *
 * Ejemplo:
 *  <CODE>
 *?     func calling[string](msg ...T) {
 *?       v := UnInterface[string](msg)
 *?	 	    parm := Option[string]("default", v)
 *?       ...
 *?     }
 *  </CODE>
 *  <ANOTHER POSIBILITY>
 *?     func calling(s ...any) {
 *? 	    s = append(s, nil, nil, nil) // inicializa los parámetros
 *?	      s[0] = map[bool]interface{}{true:  s[0], false: "cojones"}[s[0] != nil]
 *?	      s[1] = map[bool]interface{}{true:  s[1], false: 6.31}[s[1] != nil]
 *?	      s[2] = map[bool]interface{}{true:  s[2], false: 5}[s[2] != nil]
 *?	      s = s[:3] // recorta el número de parámetros
 *?       ...
 *?     }
 *  </ANOTHER POSIBILITY>
 * !@type {T comparable}
 * !@param {T} valor default, debe ser del mismo tipo que el siguiente
 * !@param {...any} valor opcional, si existe, del mismo tipo que el anterior. Debe existir.
 * !@return {T}
 */
func Opt[T comparable](_default T, v ...any) T {
	if v != nil {
		if v[0] == *new(T) {
			return _default
		}
		return v[0].(T)
	}
	return _default
}

//!-
/*******************************************************************************
 * Nos da el valor por defecto del tipo
 *
 * !@type {T comparable}
 * !@param {...T}
 * !@returns {T}
 */
func Default[T comparable](w ...T) T {
	// var (
	// 	nil T //
	// )
	return If[T](len(w) > 0)(V[T](w), *new(T)) //(Option[T](nil, w), *new(T))
	// return nil
}

/*******************************************************************************
 * Nos dice si es el valor por defecto del tipo
 *
 * !@type {T comparable}
 * !@param {T}
 * !@returns {bool}
 */
func IsDefault[T comparable](v T) bool {
	return v == Default[T]()
}

/*******************************************************************************
 * Ejecuta un decorador
 * Es parte de la otra (primera) manera de definir decoradores, la de usar Decorator[T].
 * La una parte (segunda) es la de mapas, GlobalDecorators.
 * La otra (tercera), es la de usar una concatenacion de funciones: MakeItalic(funcBB, "peluche"), que
 * es la misma que la segunda.
 *
 *<CODE>
 *? SIMPLE DECORATOR
 *? ...
 *? Decorate(func() *core.Decorator[func(func(...string) string, ...string) string] {
 *?		var (
 *?    	fn []core.Decorator[func(func(...string) string, ...string) string]
 *?   )
 *?		fn = append(fn, *(&core.Decorator[func(func(...string) string, ...string) string]{}).New(protoroos[func(...string) string]))
 *?		return fn
 *?	})(func(arg1 ...string) string {
 *?		fmt.Println("inside B with:", arg1[0])
 *?		return arg1[0]
 *?	})("sopla")
 *  ...
 *? MULTIPLE DECORATORS
 *  ...
 *?	Decorate(func() []core.Decorator[func(func(...string) string) func(...string) string] {
 *?		var (
 *?    	fn []core.Decorator[func(func(...string) string) func(...string) string]
 *?   )
 *?		fn = append(fn, *(&Decorator[func(func(...string) string) func(...string) string]{}).New(protoroos[func(...string) string]))
 *?		fn = append(fn, *(&Decorator[func(func(...string) string) func(...string) string]{}).New(MakeParagraph[func(...string) string]))
 *?		fn = append(fn, *(&Decorator[func(func(...string) string) func(...string) string]{}).New(MakeItalic[func(...string) string]))
 *?		return fn
 *?	})(funcBB)("soplapollas")
 *? ...
 *</CODE>
 *<ANOTHER CODE>
 *</ANOTHER CODE>
 * !@type {T func(func(...X) X, ...X) X, X Ordered}
 * !@param {func() *Decorator[T]}
 * !@returns {func(FuncV1X[X, X], ...X) X}
 */
func Decorate[X func(func(...T) T) func(...any) T, T comparable](deco func() []Decorator[X]) func(func(...T) T) func(...any) T {
	decor := deco()
	return func(fn func(...T) T) func(parameters ...any) T {
		return func(parameters ...any) T {
			var (
				st T
			)
			for n, v := range decor {
				if n == 0 && v.Execute {
					st = v.Decorate(fn)(parameters...) //
				} else {
					if v.Execute {
						st = v.Decorate(fn)(st) //
					}
				}
			}
			return st
		}
	}
}

/*******************************************************************************
 * Filtra los valores de value, por el tipo determinado
 * El tipo debe ser el del valor de datos a agrupar
 *
 * ! @type {T Ordered}
 * ! @param {...any}
 * ! @return {[]T}
 * ! @return {[]interface{}}
 */
func Filter[T Ordered](value ...any) (t []T, others []any) {
	var (
		n int
	)
	for n < len(value) {
		switch g := value[n].(type) {
		case T:
			t = append(t, g)
		default:
			others = append(others, g)
		}
		n++
	}
	return t, others
}

/*******************************************************************************
 * Realiza la función forEach para cada elemento del arreglo
 *
 * ! @type {T Ordered}
 * ! @param {[]T}
 * ! @param {func(T)}
 */
func ForEach[T Ordered](t []T, fn func(T)) {
	for _, v := range t {
		fn(v)
	}
}

// Variable para mantener el registro global (GlobalRegistry), si apunta a core.GlobalRegistry
// apunta al registro global del sistema apuntado por golang (deberia ser siempre asi).
// Se puede hacer que apunte a un registro particular del sistema de ficheros, p. ej. si se hace
// core.GlobalRegistry = core.Globalregistry, ó LocalRegistry ó Registry u otra cosa.
// Siempre ha de apuntar a un map[string]any
// var Registry map[string]any = make(map[string]any)
// var Registry Registry
