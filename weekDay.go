/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type (
	weekDay int // para métodos de weekDay
)

// ? @public @enum WeekDay
var WeekDay = &struct {
	Lunes, Martes, Miércoles, Jueves, Viernes, Sábado, Domingo *Const[weekDay]
}{
	Lunes:     (&Const[weekDay]{}).New(1),
	Martes:    (&Const[weekDay]{}).New(2),
	Miércoles: (&Const[weekDay]{}).New(3),
	Jueves:    (&Const[weekDay]{}).New(4),
	Viernes:   (&Const[weekDay]{}).New(5),
	Sábado:    (&Const[weekDay]{}).New(6),
	Domingo:   (&Const[weekDay]{}).New(7),
}

// interface Stringer
func (w weekDay) String() string {
	return [...]string{
		"Lunes",
		"Martes",
		"Miércoles",
		"Jueves",
		"Viernes",
		"Sabado",
		"Domingo",
	}[int(w)-1]
}
