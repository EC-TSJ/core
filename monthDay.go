/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type monthDay int

const (
	__ENERO__ monthDay = iota + 1
	__FEBRERO__
	__MARZO__
	__ABRIL__
	__MAYO__
	__JUNIO__
	__JULIO__
	__AGOSTO__
	__SEPTIEMBRE__
	__OCTUBRE__
	__NOVIEMBRE__
	__DICIEMBRE__
)

type monthday struct {
	Enero      monthDay
	Febrero    monthDay
	Marzo      monthDay
	Abril      monthDay
	Mayo       monthDay
	Junio      monthDay
	Julio      monthDay
	Agosto     monthDay
	Septiembre monthDay
	Octubre    monthDay
	Noviembre  monthDay
	Diciembre  monthDay
}

func MonthDay() *monthday {
	return &monthday{Enero: __ENERO__, Febrero: __FEBRERO__, Marzo: __MARZO__, Abril: __ABRIL__, Mayo: __MAYO__, Junio: __JUNIO__,
		Julio: __JULIO__, Agosto: __AGOSTO__, Septiembre: __SEPTIEMBRE__, Octubre: __OCTUBRE__, Noviembre: __NOVIEMBRE__, Diciembre: __DICIEMBRE__}
}

// interface Stringer
func (m monthDay) String() string {
	return [...]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}[m-1]
}
