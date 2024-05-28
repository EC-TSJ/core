/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/
/**
 * Equipación para formalización de decoradores.
 * Decorator[T] es la definición de los decoradores.
 * Decorators es una @enum que contiene las definiciones de decoradores
 * GlobalDecorators es un map que contiene decoradores
 */
/* Es la otra manera de definir decoradores */
package core

// *****************************************************************************
// Definición del decorador a emplear
// *****************************************************************************
// ? @public @class Decorator[T]
// *****************************************************************************
type Decorator[T any] struct {
	// /**
	//  *! @mixin
	//  */*eventer.MEvent
	Class    ClassType
	Decorate T    //Tipo de función que decora la función --> el decorador
	Execute  bool // ejecuta la función de decoración
}

/**
 *? @static @constructor @New
 *
 * Crea el Decorador
 *
 *! @param {T} function la fn que crea el decorador --> el decorador, en una palabra
 *! @return {*Decorator[T]}
 */
func (this *Decorator[T]) New(function T) *Decorator[T] {
	return &Decorator[T]{Class: Class.Normal, Decorate: function, Execute: true}
}

/**
 * Obtiene la ejecución del decorador
 *
 *! @return {bool}
 *? core.Method.Normal
 */
func (this *Decorator[T]) Get() bool {
	flag := this.Execute
	this.Execute = false
	return flag
}

/**
 * Setea la ejecucicón del decorador
 *
 *! @param {bool}
 *? core.Method.Normal
 */
func (this *Decorator[T]) Set(x bool) {
	this.Execute = x
}

/**-----------------------------------------------------------------------------
 *? @endclass
 *------------------------------------------------------------------------------
 */
