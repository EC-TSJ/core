package core

/*******************************************************************************
 * Types[T ordered]. @Abstract Class de la que dependen las demás.
 *? @public @abstract @class types[T] @implements IAccesors[T]
 ******************************************************************************/
type types[T Ordered] struct {
	/**
		 *! @extends
	   *? @@AbstractError
	*/*Decorator[func(func(...any) T) func(ClassType, Object_Of) T]
	/**
	 *! @implements
	 */IAccesors[T]
	Class      ClassType
	__assigned bool // solo para ReadOnly
	__key      T    // valor
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @compose
 */
func (*types[T]) compose(val T, cl ClassType) *types[T] {
	return &types[T]{
		IAccesors:  nil,
		Class:      cl,
		__key:      val,
		__assigned: true,
		Decorator:  (&Decorator[func(func(...any) T) func(ClassType, Object_Of) T]{}).New(AbstractError),
	}
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @new
 */
//? @AbstractError
func (*types[T]) New() {
	AbstractError := (&Decorator[func(func(...any) T) func(ClassType, Object_Of) T]{}).New(AbstractError)
	AbstractError.Decorate(func(...any) T { return Val(-1).(T) })(Class.Abstract, ObjectOf.Class)
}

/*******************************************************************************
 *? @virtual
 *? @override
 */
//? @AbstractError
func (this *types[T]) Get() T {
	return this.Decorate(func(...any) T { return this.__key })(this.Class, ObjectOf.Method)
}

/*******************************************************************************
 *? @virtual
 *? @override
 */
//? @AbstractError
func (this *types[T]) Set(value T) {
	this.Decorate(func(...any) T { return Val(-1).(T) })(this.Class, ObjectOf.Method)
}

/*******************************************************************************
 * Nos da un valor de tipo ReadOnly.
 * Se asigna un valor de inicio pero ya no se puede modificar
 *******************************************************************************
 *? @public @class ReadOnly[T] @extends types[T]
 ******************************************************************************/
type ReadOnly[T Ordered] struct {
	/**
	 *! @extends
	 */types[T]
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @New
 */
func (this *ReadOnly[T]) New(value T) *ReadOnly[T] {
	var (
		nil T
		rd  *ReadOnly[T] = this
	)
	if !this.__assigned {
		rd = &ReadOnly[T]{*(&types[T]{}).compose(value, Class.Derived)}
		rd.IAccesors = rd
		if value == nil {
			rd.__assigned = false
		}
	}
	return rd
}

/*******************************************************************************
 *? @override
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
 *? @public @class Const[T] @extends types[T]
 ******************************************************************************/
type Const[T Ordered] struct {
	/**
	 *! @extends
	 */types[T]
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @New
 */
func (this *Const[T]) New(val T) *Const[T] {
	var (
		_const *Const[T] = this
	)
	if !this.__assigned {
		_const = &Const[T]{*(&types[T]{}).compose(val, Class.Derived)}
		_const.IAccesors = _const
	}
	return _const
}

/*******************************************************************************
 * Nos da un valor de tipo Static. Del tipo static de C, no de las clases de
 * instancia/valor.
 * Los límites en donde se mueve son Upper y Lower, de Java, p. ej.
 *******************************************************************************
 *? @public @class Static[T] @extends types[T]
 ******************************************************************************/
type Static[T Number] struct {
	/**
	 *! @extends
	 */types[T]
	Lower int // lower index for print
	Upper int // upeer index for print
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @New
 */
func (this *Static[T]) New(val T) *Static[T] {
	st := &Static[T]{*(&types[T]{}).compose(val, Class.Derived), 0, 9223372036854775806}
	st.IAccesors = st
	return st
}

/*******************************************************************************
 *? @override
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
// ? @type prop
type prop byte

const (
	// Lector de propiedades habilitado
	T_PROPGET prop = 0b0001
	// Escritor de propiedades habilitado
	T_PROPSET prop = 0b0010
	// valor cero
	zero prop = 0
)

/*******************************************************************************
 * Nos da un dato. Del que sea s/ T.
 * Se puede usar como las propiedades Get/Set de C#, en
 * tal caso, se usan las constantes T_PROPGET y  T_PROPSET
 * para controlar cual son las funciones que se asignan: Get
 * ó Set.
 * Se usa como el segundo parámetro de New(Value, [Propertys]),
 * ó directamente a través de la propiedad prop. Así:
 * InstanceVar(propertyInstance[type]).prop = T_PROPGET | T_PROPSET.
 *******************************************************************************
 *? @public @class Property[T] @extends types[T]
 ******************************************************************************/
type Property[T Ordered] struct {
	/**
	 *! @extends
	 */types[T]
	prop prop
}

/*******************************************************************************
 *? @static
 *? @constructor
 *? @New
 */
func (this *Property[T]) New(value T, t ...prop) *Property[T] {
	var (
		_prop *Property[T] = this
		z     prop
	)
	if !this.__assigned {
		z = (&Ovl2[prop]{}).New().ArgOptional(zero, t) // optional[prop](zero, t...)
		_prop = &Property[T]{types: *(&types[T]{}).compose(value, Class.Derived), prop: If[prop](t == nil)(T_PROPGET|T_PROPSET, z)}
		_prop.IAccesors = _prop
	}
	return _prop
}

/*******************************************************************************
 *? @override
 */
func (this *Property[T]) Get() /**@override*/ T {
	if this.prop&T_PROPGET == T_PROPGET {
		return this.__key
	}
	return *new(T) //Default[T]()
}

/*******************************************************************************
 *? @override
 */
func (this *Property[T]) Set(value T) /**@override*/ {
	if this.prop&T_PROPSET == T_PROPSET {
		this.__key = value
	}
}

// func optional[T Ordered](_default T, _optional ...T) T {
// 	if len(_optional) > 0 {
// 		return _optional[0]
// 	}
// 	return _default
// }
