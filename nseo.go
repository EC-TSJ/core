/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type (
	NSEO   int
	nseoSt struct {
		// Norte
		Norte NSEO
		// Sur
		Sur NSEO
		// Este
		Este NSEO
		// Oeste
		Oeste NSEO
	}
	Enseo = nseoSt
)

const (
	__NORTE__ NSEO = iota + 1
	__SUR__
	__ESTE__
	__OESTE__
)

// enum NSEO
var ENSEO *Enseo = &nseoSt{Norte: __NORTE__, Sur: __SUR__, Este: __ESTE__, Oeste: __OESTE__}

// Interface Stringer
func (d NSEO) String() string {
	return [...]string{"Norte", "Sur", "Este", "Oeste"}[d-1]
}
