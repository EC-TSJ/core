package core

import (
	"errors"
)

/*******************************************************************************
 * ATypes[T ordered]. @Abstract Class de la que dependen las demás.
 *? @public @abstract @class Types[T] @implements IGetter
 ******************************************************************************/
type Types[T Ordered] struct {
	/**
	 *! @implements
	 */IAccesors IAccesors[T]
	Class        ClassType
	__assigned   bool // solo para ReadOnly
	__key        T    // valor
}

/**************************************************
 * decorators: __ending && __panicking
 *************************************************/
//? @decorator __ending[T](*Types[T])(classMethods)
func __ending[T Ordered](g *Types[T]) func(ClassMethods) {
	return func(vis ClassMethods) {
		if g.Class == Class.Abstract {
			var (
				s string = [...]string{"Class", "Method"}[vis]
				e error  = errors.New("ReferenceError: Abstract " + s + ": Don't Execute.")
			)
			Exception(nil)(func() { Throw(e) })
		}
	}
}

// ? @decorator  __panicking[T](*Types[T])()
//
// ? @__ending(g)(ClassMethod)
func __panicking[T Ordered](this *Types[T]) func() {
	return func() {
		__ending(this)(ClassMethod.Method)
	}
}

/*******************************************************************************
 *? @media @constructor
 */
func (*Types[T]) compose(val T, cl ClassType) *Types[T] {
	return &Types[T]{IAccesors: nil, Class: cl, __key: val, __assigned: true}
}

/*******************************************************************************
 *? @static @constructor @new
 */
//? @__ending(this)(ClassMethod)
func (this *Types[T]) New() {
	__ending(this)(ClassMethod.Class)
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
//? @__panicking(g)()
func (this *Types[T]) Get() /**@virtual*/ T {
	__panicking(this)()
	return this.__key
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
//? @__panicking(g)()
func (this *Types[T]) Set(value T) /**@virtual*/ {
	__panicking(this)()
}

/*******************************************************************************
 * Nos da un valor de tipo ReadOnly.
 * Se asigna un valor de inicio pero ya no se puede modificar
 *******************************************************************************
 *? @public @class ReadOnly[T] @extends Types[T]
 ******************************************************************************/
type ReadOnly[T Ordered] struct {
	/**
	 *! @extends
	 */Types[T]
}

/*******************************************************************************
 *? @static @constructor @New
 */
func (this *ReadOnly[T]) New(value T) *ReadOnly[T] {
	var (
		nil T
		rd  *ReadOnly[T] = this
	)
	if !this.__assigned {
		rd = &ReadOnly[T]{*(&Types[T]{}).compose(value, Class.Derived)}
		rd.IAccesors = rd
		if value == nil {
			rd.__assigned = false
		}
	}
	return rd
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
func (this *ReadOnly[T]) Set(value T) /**@override*/ {
	if !this.__assigned {
		this.__key = value
		this.__assigned = true
	}
}

/*******************************************************************************
 * Nos da un valor de tipo Const.
 *******************************************************************************
 *? @public @class Const[T] @extends Types[T]
 ******************************************************************************/
type Const[T Ordered] struct {
	/**
	 *! @extends
	 */Types[T]
}

/*******************************************************************************
 *? @static @constructor @New
 */
func (this *Const[T]) New(val T) *Const[T] {
	var (
		_const *Const[T] = this
	)
	if !this.__assigned {
		_const = &Const[T]{*(&Types[T]{}).compose(val, Class.Derived)}
		_const.IAccesors = _const
	}
	return _const
}

/*******************************************************************************
 * Nos da un valor de tipo Static. Del tipo static de C, no de las clases de
 * instancia/valor.
 * Los límites en donde se mueve son Upper y Lower, de Java, p. ej.
 *******************************************************************************
 *? @public @class Static[T] @extends Types[T]
 ******************************************************************************/
type Static[T Number] struct {
	/**
	 *! @extends
	 */Types[T]
	Lower int // lower index for print
	Upper int // upeer index for print
}

/*******************************************************************************
 *? @static @constructor @New
 */
func (this *Static[T]) New(val T) *Static[T] {
	st := &Static[T]{*(&Types[T]{}).compose(val, Class.Derived), 0, 9223372036854775806}
	st.IAccesors = st
	return st
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
func (this *Static[T]) Get() /**@override*/ T {
	if this.__key > T(this.Upper) {
		this.__key = T(this.Lower)
	}
	pd := this.__key
	this.__key += T(1)
	return pd
}

// ------------------------------------------------------------------------------
// ? @type Prop
type Prop byte

const (
	// Lector de propiedades habilitado
	T_PROPGET Prop = 0b0001
	// Escritor de propiedades habilitado
	T_PROPSET Prop = 0b0010
	// valor cero
	zero Prop = 0
)

/*******************************************************************************
 * Nos da un dato. Del que sea s/ T.
 * Se puede usar como las propiedades Get/Set de C#, en
 * tal caso, se usan las constantes T_PROPGET y  T_PROPSET
 * para controlar cual son las funciones que se asignan: Get
 * ó Set.
 * Se usa como el segundo parámetro de New(Value, [Propertys]),
 * ó directamente a través de la propiedad Prop. Así:
 * InstanceVar(propertyInstance[type]).Prop = T_PROPGET | T_PROPSET.
 *******************************************************************************
 *? @public @class Property[T] @extends Types[T]
 ******************************************************************************/
type Property[T Ordered] struct {
	/**
	 *! @extends
	 */Types[T]
	Prop Prop
}

/*******************************************************************************
 *? @static @constructor @New
 */
func (this *Property[T]) New(value T, t ...Prop) *Property[T] {
	var (
		_prop *Property[T] = this
		z     Prop
	)
	if !this.__assigned {
		z = (&Ovl2[Prop]{}).New().ArgOptional(zero, t) // optional[Prop](zero, t...)
		_prop = &Property[T]{Types: *(&Types[T]{}).compose(value, Class.Derived), Prop: If[Prop](t == nil)(T_PROPGET|T_PROPSET, z)}
		_prop.IAccesors = _prop
	}
	return _prop
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
func (this *Property[T]) Get() /**@override*/ T {
	if this.Prop&T_PROPGET == T_PROPGET {
		return this.__key
	}
	return *new(T) //Default[T]()
}

/*******************************************************************************
 *? @override
 *? core.Method.Override
 */
func (this *Property[T]) Set(value T) /**@override*/ {
	if this.Prop&T_PROPSET == T_PROPSET {
		this.__key = value
	}
}

// func optional[T Ordered](_default T, _optional ...T) T {
// 	if len(_optional) > 0 {
// 		return _optional[0]
// 	}
// 	return _default
// }
