package core

import "reflect"

type (
	/**
	 *? @public @abstract @interface IPair[T1, T2]
	 */
	IPair[T1, T2 any] interface {
		/**
		 *? @abstract @virtual
		 */
		GetT1() /** @virtual */ T1
		/**
		 *? @abstract @virtual
		 */
		GetT2() /** @virtual */ T2
	}

	/*******************************************************************************
	 * Tipo APair
	 *! @abstract @class
	 *! @type {T1 any}
	 *! @type {T2 anyº}
	 *
	 *? @public @abstract @class APair[T1, T2]
	 */
	APair[T1, T2 any] struct {
		IPair[T1, T2]
		Class ClassType
		T1    T1
		T2    T2
	}

	/*******************************************************************************
	 * Tipo Pair
	 *! @Class
	 *! @type {T1 any}
	 *! @type {T2 anyº}
	 *
	 *? @public @class Pair[T1, T2]
	 */
	Pair[T1, T2 any] struct {
		*APair[T1, T2]
	}
)

/*******************************************************************************
 *  Crea un nuevo Pair[T,T]
 *! @param {T1}
 *! @param {T2}
 *! @returns {*Pair[T1, T2]}
 *
 *? @static @constructor @New
 */
func (*Pair[T1, T2]) New(t1 T1, t2 T2) *Pair[T1, T2] {
	pair := &Pair[T1, T2]{APair: &APair[T1, T2]{IPair: nil, Class: Class.Normal, T1: t1, T2: t2}}
	pair.IPair = pair
	return pair
}

/*******************************************************************************
 *  Obtiene el parámetro T1
 *! @returns {T1}
 *
 *? @override
 *? core.Method.Override
 */
func (this *Pair[T1, T2]) GetT1() /** @override */ T1 {
	return this.T1
}

/*******************************************************************************
 *  Obtiene el parámetro T2
 *! @returns {T2}
 *
 *? @override
 *? core.Method.Override
 */
func (this *Pair[T1, T2]) GetT2() /** @override */ T2 {
	return this.T2
}

////////////////////////////////////////////////////////////////////////////////
//! @Stringer
/*******************************************************************************
 *  Obtiene el nombre de la clase
 *! @returns {string}
 *
 *? @core.Method.Normal
 */
func (this *APair[T1, T2]) String() string {
	return If[string](reflect.TypeOf(this).Kind() == reflect.Pointer)("*"+reflect.TypeOf(this).Elem().Name(), reflect.TypeOf(this).Name())
}
