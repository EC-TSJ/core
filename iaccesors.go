package core

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface IGetter
 ******************************************************************************/
type IGetter[T Ordered] interface {
	/**
	 *? @abstract @virtual
	 */
	Get() /**@virtual*/ T
}

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface ISetter
 ******************************************************************************/
type ISetter[T Ordered] interface {
	/**
	 *? @abstract @virtual
	 */
	Set(T) /**@virtual*/
}

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface IAccesors
 ******************************************************************************/
type IAccesors[T Ordered] interface {
	IGetter[T]
	ISetter[T]
}
