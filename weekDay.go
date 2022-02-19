/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type (
	WeekDay int
	weekday struct {
		// Lunes
		Lunes WeekDay
		// Martes
		Martes WeekDay
		// Miércoles
		Miércoles WeekDay
		// Jueves
		Jueves WeekDay
		// Viernes
		Viernes WeekDay
		// Sábado
		Sábado WeekDay
		// Domingo
		Domingo WeekDay
	}
	EWeekday = weekday
)

const (
	__LUNES__ WeekDay = iota + 1
	__MARTES__
	__MIERCOLES__
	__JUEVES__
	__VIERNES__
	__SABADO__
	__DOMINGO__
)

// enum WeekDay
var EWeekDay *EWeekday = &weekday{Lunes: __LUNES__, Martes: __MARTES__, Miércoles: __MIERCOLES__, Jueves: __JUEVES__,
	Viernes: __VIERNES__, Sábado: __SABADO__, Domingo: __DOMINGO__}

// interface Stringer
func (w WeekDay) String() string {
	return [...]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sabado", "Domingo"}[w-1]
}
