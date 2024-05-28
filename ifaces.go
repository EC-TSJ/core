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
