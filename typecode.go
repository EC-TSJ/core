/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

// TypeCode del dato
type TypeCode int

const (
	_EMPTY    TypeCode = iota // Null reference
	_OBJECT                   // Instance that isn't a value
	_DBNULL                   // Database null value
	_BOOLEAN                  // Boolean
	_CHAR                     // Unicode character
	_INT8                    // Signed 8-bit integer
	_UINT8
	_INT16                    // Signed 16-bit integer
	_UINT16                   // Unsigned 16-bit integer
	_INT32                   // Signed 32-bit integer
	_UINT32                   // Unsigned 32-bit integer
	_INT64                    // Signed 64-bit integer
	_UINT64                   // Unsigned 64-bit integer
	_FLOAT32                   // IEEE 32-bit float
	_FLOAT64                   // IEEE 64-bit double
	_COMPLEX64                  // Decimal
	_COMPLEX128
	_DATETIME                 // DateTime
	_STRING                 // Unicode character string
)
