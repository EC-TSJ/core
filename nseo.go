/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type nseo int

const (
	__NORTE__ nseo = iota + 1
	__SUR__
	__ESTE__
	__OESTE__
)

type nseoSt struct {
	Norte nseo
	Sur   nseo
	Este  nseo
	Oeste nseo
}

func NSEO() *nseoSt {
	return &nseoSt{Norte: __NORTE__, Sur: __SUR__, Este: __ESTE__, Oeste: __OESTE__}
}

// Interface Stringer
func (d nseo) String() string {
	return [...]string{"Norte", "Sur", "Este", "Oeste"}[d-1]
}
