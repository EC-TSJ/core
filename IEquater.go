/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type IEquater interface {
	Equals(int) bool
}

type IEquaterIf interface {
	Equals(T) bool
}

type IEquaterS interface {
	Equals(string) bool
}
