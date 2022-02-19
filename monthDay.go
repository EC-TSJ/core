/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type (
	MonthDay int
	monthday struct {
		// Enero
		Enero MonthDay
		// Ferbrero
		Febrero MonthDay
		// Marzo
		Marzo MonthDay
		// Abril
		Abril MonthDay
		// Mayo
		Mayo MonthDay
		// Junio
		Junio MonthDay
		// Julio
		Julio MonthDay
		// Agosto
		Agosto MonthDay
		// Septiembre
		Septiembre MonthDay
		// Octubre
		Octubre MonthDay
		// Noviembre
		Noviembre MonthDay
		// Diciembre
		Diciembre MonthDay
	}
	EMonthday = monthday
)

const (
	__ENERO__ MonthDay = iota + 1
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

// enum MonthDay
var EMonthDay *EMonthday = &monthday{Enero: __ENERO__, Febrero: __FEBRERO__, Marzo: __MARZO__, Abril: __ABRIL__, Mayo: __MAYO__, Junio: __JUNIO__,
	Julio: __JULIO__, Agosto: __AGOSTO__, Septiembre: __SEPTIEMBRE__, Octubre: __OCTUBRE__, Noviembre: __NOVIEMBRE__, Diciembre: __DICIEMBRE__}

// interface Stringer
func (m MonthDay) String() string {
	return [...]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}[m-1]
}
