/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

// interface IComparer
type IComparer interface {
	Compare(int) int
}

// interface IComparerIf
type IComparerIf interface {
	Compare(T) int
}

// interface IComparerS
type IComparerS interface {
	Compare(string) string
}

// interface IComparerS
type IComparerIfS interface {
	Compare(T) string
}
