/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

/**************************************************
 * for use instead of cmp.IComparer[T](x, y T) int
 * ? @public  @interface IComparer[T]
 ***************************************************/
type IComparer[T Ordered] interface {
	Compare(T) int
}

/**************************************************
 * ? @public  @interface IEquater[T]
 ***************************************************/
// ? @public @interface IEquater
type IEquater[T comparable] interface {
	Equal(T) bool
}

// ? @public @interface ICounter
type ICounter interface {
	Count() int
}

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface IGetter
 ******************************************************************************/
type IGetter[T Ordered] interface {
	/**
	 *? @abstract
	 *? @virtual
	 */
	Get() T
}

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface ISetter
 ******************************************************************************/
type ISetter[T Ordered] interface {
	/**
	 *? @abstract
	 *? @virtual
	 */
	Set(T)
}

/*******************************************************************************
 *  for Get/Set functions
 *? @public @interface IAccesors
 ******************************************************************************/
type IAccesors[T Ordered] interface {
	IGetter[T]
	ISetter[T]
}
