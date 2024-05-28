/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

type (
	monthDay int // Para métodos de monthDay
)

// ? @public @enum MonthDay
var MonthDay = &struct {
	Enero, Febrero, Marzo, Abril, Mayo, Junio, Julio, Agosto, Septiembre, Octubre, Noviembre, Diciembre *Const[monthDay]
}{
	Enero:      (&Const[monthDay]{}).New(1),
	Febrero:    (&Const[monthDay]{}).New(2),
	Marzo:      (&Const[monthDay]{}).New(3),
	Abril:      (&Const[monthDay]{}).New(4),
	Mayo:       (&Const[monthDay]{}).New(5),
	Junio:      (&Const[monthDay]{}).New(6),
	Julio:      (&Const[monthDay]{}).New(7),
	Agosto:     (&Const[monthDay]{}).New(8),
	Septiembre: (&Const[monthDay]{}).New(9),
	Octubre:    (&Const[monthDay]{}).New(10),
	Noviembre:  (&Const[monthDay]{}).New(11),
	Diciembre:  (&Const[monthDay]{}).New(12),
}

// interface Stringer
func (m monthDay) String() string {
	return [...]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}[int(m)-1]
}
