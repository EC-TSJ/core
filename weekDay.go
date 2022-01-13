/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type weekDay int

const (
	__LUNES__ weekDay = iota + 1
	__MARTES__
	__MIERCOLES__
	__JUEVES__
	__VIERNES__
	__SABADO__
	__DOMINGO__
)

type weekday struct {
	Lunes     weekDay
	Martes    weekDay
	Miércoles weekDay
	Jueves    weekDay
	Viernes   weekDay
	Sábado    weekDay
	Domingo   weekDay
}

func WeekDay() *weekday {
	return &weekday{Lunes: __LUNES__, Martes: __MARTES__, Miércoles: __MIERCOLES__, Jueves: __JUEVES__,
		Viernes: __VIERNES__, Sábado: __SABADO__, Domingo: __DOMINGO__}
}

// interface Stringer
func (w weekDay) String() string {
	return [...]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sabado", "Domingo"}[w-1]
}
