/************************************/
/*  , %T% %S%), %J% <$B$> <$1.00$>   */
/*  ,%W% 30-09-1991 )               */
/*  ,%X%            )               */
/*  ,%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"fmt"
	"strings"
)

const ()

var (
	szA = "Informational"
	szB = "Successful"
	szC = "Redirection"
	szD = "ClientError"
	szE = "ServerError"
	//Errores de HTTPError
	httpCodes = map[string]*Pair{
		// 1xx - Informational
		strings.ToUpper("Continue"):           MakePair(100, szA),
		strings.ToUpper("SwitchingProtocols"): MakePair(101, szA),
		strings.ToUpper("Processing"):         MakePair(102, szA),
		strings.ToUpper("EarlyHints"):         MakePair(103, szA),
		// 2xx: Successful
		strings.ToUpper("OK"):                          MakePair(200, szB),
		strings.ToUpper("Created"):                     MakePair(201, szB),
		strings.ToUpper("Accepted"):                    MakePair(202, szB),
		strings.ToUpper("NonAuthoritativeInformation"): MakePair(203, szB),
		strings.ToUpper("NoContent"):                   MakePair(204, szB),
		strings.ToUpper("ResetContent"):                MakePair(205, szB),
		strings.ToUpper("PartialContent"):              MakePair(206, szB),
		strings.ToUpper("MultiStatus"):                 MakePair(207, szB),
		strings.ToUpper("AlreadyReported"):             MakePair(208, szB),
		strings.ToUpper("ImUsed"):                      MakePair(226, szB),
		// 3xx: Redirection
		strings.ToUpper("MultipleChoices"):   MakePair(300, szC),
		strings.ToUpper("MovedPermanently"):  MakePair(301, szC),
		strings.ToUpper("Found"):             MakePair(302, szC),
		strings.ToUpper("SeeOther"):          MakePair(303, szC),
		strings.ToUpper("NotModified"):       MakePair(304, szC),
		strings.ToUpper("UseProxy"):          MakePair(305, szC),
		strings.ToUpper("SwitchProxy"):       MakePair(306, szC),
		strings.ToUpper("TemporaryRedirect"): MakePair(307, szC),
		strings.ToUpper("PermanentRedirect"): MakePair(308, szC),
		// 4xx: Client Error
		strings.ToUpper("BadRequest"):                  MakePair(400, szD),
		strings.ToUpper("Unauthorized"):                MakePair(401, szD),
		strings.ToUpper("PaymentRequired"):             MakePair(402, szD),
		strings.ToUpper("Forbidden"):                   MakePair(403, szD),
		strings.ToUpper("NotFound"):                    MakePair(404, szD),
		strings.ToUpper("MethodNotAllowed"):            MakePair(405, szD),
		strings.ToUpper("NotAcceptable"):               MakePair(406, szD),
		strings.ToUpper("ProxyAuthenticationRequired"): MakePair(407, szD),
		strings.ToUpper("RequestTimeout"):              MakePair(408, szD),
		strings.ToUpper("Conflict"):                    MakePair(409, szD),
		strings.ToUpper("Gone"):                        MakePair(410, szD),
		strings.ToUpper("LengthRequired"):              MakePair(411, szD),
		strings.ToUpper("PreconditionFailed"):          MakePair(412, szD),
		strings.ToUpper("PayloadTooLarge"):             MakePair(413, szD),
		strings.ToUpper("URITooLong"):                  MakePair(414, szD),
		strings.ToUpper("UnsupportedMediaType"):        MakePair(415, szD),
		strings.ToUpper("RangeNotSatisfiable"):         MakePair(416, szD),
		strings.ToUpper("ExpectationFailed"):           MakePair(417, szD),
		strings.ToUpper("ImATeapot"):                   MakePair(418, szD),
		strings.ToUpper("MisdirectedRequest"):          MakePair(421, szD),
		strings.ToUpper("UnprocessableEntity"):         MakePair(422, szD),
		strings.ToUpper("Locked"):                      MakePair(423, szD),
		strings.ToUpper("FailedDependency"):            MakePair(424, szD),
		strings.ToUpper("TooEarly"):                    MakePair(425, szD),
		strings.ToUpper("UpgradeRequired"):             MakePair(426, szD),
		strings.ToUpper("PreconditionRequired"):        MakePair(428, szD),
		strings.ToUpper("TooManyRequests"):             MakePair(429, szD),
		strings.ToUpper("RequestHeaderFieldsTooLarge"): MakePair(431, szD),
		strings.ToUpper("UnavailableForLegalReasons"):  MakePair(451, szD),
		// 5xx: Server Error
		strings.ToUpper("InternalServerError"):           MakePair(500, szE),
		strings.ToUpper("NotImplemented"):                MakePair(501, szE),
		strings.ToUpper("BadGateway"):                    MakePair(502, szE),
		strings.ToUpper("ServiceUnavailable"):            MakePair(503, szE),
		strings.ToUpper("GatewayTimeout"):                MakePair(504, szE),
		strings.ToUpper("HTTPVersionNotSupported"):       MakePair(505, szE),
		strings.ToUpper("VariantAlsoNegotiates"):         MakePair(506, szE),
		strings.ToUpper("InsufficientStorage"):           MakePair(507, szE),
		strings.ToUpper("LoopDetected"):                  MakePair(508, szE),
		strings.ToUpper("NotExtended"):                   MakePair(510, szE),
		strings.ToUpper("NetworkAuthenticationRequired"): MakePair(511, szE),
	}

	szDBA = "ArgumentError"
	szDBB = "DBError"
	szDBC = "InvalidRequestError"
	szDBD = "NoReferenceError"
	szDBE = "DBAPIError"
	szDBF = "DatabaseError"
	//Errores de DBError
	dbCodes = map[string]*Pair{
		strings.ToUpper(szDBA):                         MakePair(szDBB, 0x001),
		strings.ToUpper("ObjectNotExecutableError"):    MakePair(szDBA, 0x002),
		strings.ToUpper("NoSuchModuleError"):           MakePair(szDBA, 0x003),
		strings.ToUpper("NoForeignKeysError"):          MakePair(szDBA, 0x004),
		strings.ToUpper("AmbiguousForeignKeysError"):   MakePair(szDBA, 0x005),
		strings.ToUpper("CircularDependencyError"):     MakePair(szDBB, 0x006),
		strings.ToUpper("CompileError"):                MakePair(szDBB, 0x007),
		strings.ToUpper("UnsupportedCompilationError"): MakePair("CompileError", 0x008),
		strings.ToUpper("IdentifierError"):             MakePair(szDBB, 0x009),
		strings.ToUpper("DisconnectionError"):          MakePair(szDBB, 0x01A),
		strings.ToUpper("InvalidatePoolError"):         MakePair("DisconnectionError", 0x01B),
		strings.ToUpper("TimeoutError"):                MakePair(szDBB, 0x01C),
		strings.ToUpper(szDBC):                         MakePair(szDBB, 0x01D),
		strings.ToUpper("NoInspectionAvailableError"):  MakePair(szDBC, 0x01E),
		strings.ToUpper("ResourceClosedError"):         MakePair(szDBC, 0x01F),
		strings.ToUpper("NoSuchColumnError"):           MakePair(szDBC, 0x020),
		strings.ToUpper(szDBD):                         MakePair(szDBC, 0x021),
		strings.ToUpper("NoReferencedTableError"):      MakePair(szDBD, 0x022),
		strings.ToUpper("NoReferencedColumnError"):     MakePair(szDBD, 0x023),
		strings.ToUpper("NoSuchTableError"):            MakePair(szDBC, 0x24),
		strings.ToUpper("UnreflectableTableError"):     MakePair(szDBC, 0x25),
		strings.ToUpper("UnboundExecutionError"):       MakePair(szDBC, 0x26),
		strings.ToUpper("StatementError"):              MakePair(szDBB, 0x27),
		strings.ToUpper(szDBE):                         MakePair("StatementError", 0x28),
		strings.ToUpper("InterfaceError"):              MakePair(szDBE, 0x29),
		strings.ToUpper(szDBF):                         MakePair(szDBE, 0x2A),
		strings.ToUpper("DataError"):                   MakePair(szDBF, 0x2B),
		strings.ToUpper("OperationalError"):            MakePair(szDBF, 0x2C),
		strings.ToUpper("IntegrityError"):              MakePair(szDBF, 0x2D),
		strings.ToUpper("InternalError"):               MakePair(szDBF, 0x2E),
		strings.ToUpper("ProgrammingError"):            MakePair(szDBF, 0x2F),
		strings.ToUpper("NotSupportedError"):           MakePair(szDBF, 0x30),
	}

	szFIOA = "OSError"
	szFIOB = "ConnectionError"
	//Errores de FileIOError
	fileIOCodes = map[string]*Pair{
		strings.ToUpper(szFIOA):                   MakePair("FileIOError", 0x100),
		strings.ToUpper("BlockingIOError"):        MakePair(szFIOA, 0x101),
		strings.ToUpper("ChildProcessError"):      MakePair(szFIOA, 0x102),
		strings.ToUpper(szFIOB):                   MakePair(szFIOA, 0x103),
		strings.ToUpper("BrokenPipeError"):        MakePair(szFIOB, 0x104),
		strings.ToUpper("ConnectionAbortedError"): MakePair(szFIOB, 0x105),
		strings.ToUpper("ConnectionRefusedError"): MakePair(szFIOB, 0x106),
		strings.ToUpper("ConnectionResetError"):   MakePair(szFIOB, 0x107),
		strings.ToUpper("FileExistsError"):        MakePair(szFIOA, 0x108),
		strings.ToUpper("FileNotFoundError"):      MakePair(szFIOA, 0x109),
		strings.ToUpper("InterruptedError"):       MakePair(szFIOA, 0x10A),
		strings.ToUpper("IsADirectoryError"):      MakePair(szFIOA, 0x10B),
		strings.ToUpper("NotADirectoryError"):     MakePair(szFIOA, 0x10C),
		strings.ToUpper("PermissionError"):        MakePair(szFIOA, 0x10D),
		strings.ToUpper("ProcessLookupError"):     MakePair(szFIOA, 0x10E),
		strings.ToUpper("TimeoutError"):           MakePair(szFIOA, 0x10F),
	}

	szBaseA = "Error"
	szBaseB = "ArithmeticError"
	szBaseC = "UnicodeError"
	szBaseD = "LookupError"
	szBaseE = "RuntimeError"
	szBaseF = "BaseError"
	//Errores de Error
	baseErrorCodes = map[string]*Pair{
		strings.ToUpper(szBaseB):                   MakePair(szBaseA, 0x200),
		strings.ToUpper("FloatingPointError"):      MakePair(szBaseB, 0x201),
		strings.ToUpper("OverflowError"):           MakePair(szBaseB, 0x202),
		strings.ToUpper("ZeroDivisionError"):       MakePair(szBaseB, 0x203),
		strings.ToUpper("AssertionError"):          MakePair(szBaseA, 0x204),
		strings.ToUpper("AttributeError"):          MakePair(szBaseA, 0x205),
		strings.ToUpper("BufferError"):             MakePair(szBaseA, 0x206),
		strings.ToUpper("EOFError"):                MakePair(szBaseA, 0x207),
		strings.ToUpper("ImportError"):             MakePair(szBaseA, 0x208),
		strings.ToUpper("ModuleNotFoundError"):     MakePair("ImportError", 0x209),
		strings.ToUpper(szBaseD):                   MakePair(szBaseA, 0x20A),
		strings.ToUpper("IndexError"):              MakePair(szBaseD, 0x20B),
		strings.ToUpper("KeyError"):                MakePair(szBaseD, 0x20C),
		strings.ToUpper("MemoryError"):             MakePair(szBaseA, 0x20D),
		strings.ToUpper("NameError"):               MakePair(szBaseA, 0x20E),
		strings.ToUpper("UnboundLocalError"):       MakePair("NameError", 0x20F),
		strings.ToUpper("ReferenceError"):          MakePair(szBaseA, 0x210),
		strings.ToUpper(szBaseE):                   MakePair(szBaseA, 0x211),
		strings.ToUpper("NotImplementedError"):     MakePair(szBaseE, 0x212),
		strings.ToUpper("RecursionError"):          MakePair(szBaseE, 0x213),
		strings.ToUpper("StopIterationError"):      MakePair(szBaseA, 0x214),
		strings.ToUpper("StopAsyncIterationError"): MakePair(szBaseA, 0x215),
		strings.ToUpper("SyntaxError"):             MakePair(szBaseA, 0x216),
		strings.ToUpper("IndentationError"):        MakePair("SyntaxError", 0x2217),
		strings.ToUpper("TabError"):                MakePair("IndentationError", 0x218),
		strings.ToUpper("SystemError"):             MakePair(szBaseA, 0x219),
		strings.ToUpper("TypeError"):               MakePair(szBaseA, 0x22A),
		strings.ToUpper("ValueError"):              MakePair(szBaseA, 0x22B),
		strings.ToUpper(szBaseC):                   MakePair("ValueError", 0x22C),
		strings.ToUpper("UnicodeDecodeError"):      MakePair(szBaseC, 0x22D),
		strings.ToUpper("UnicodeEncodeError"):      MakePair(szBaseC, 0x22E),
		strings.ToUpper("UnicodeTranslateError"):   MakePair(szBaseC, 0x22F),
		strings.ToUpper("GeneratorExit"):           MakePair(szBaseF, 0x233),
		strings.ToUpper("KeyboardInterrupt"):       MakePair(szBaseF, 0x234),
		strings.ToUpper("SystemExit"):              MakePair(szBaseF, 0x235),
		strings.ToUpper("ParameterError"):          MakePair(szBaseA, 0x236),
		strings.ToUpper("Correct"):                 MakePair("OperationValidated", 0x300),
	}

	szW = "Warning"
	//Errores de Warning
	warningCodes = map[string]*Pair{
		strings.ToUpper(szW):                         MakePair(szW, 0x300),
		strings.ToUpper("BytesWarning"):              MakePair(szW, 0x301),
		strings.ToUpper("DeprecationWarning"):        MakePair(szW, 0x302),
		strings.ToUpper("FutureWarning"):             MakePair(szW, 0x303),
		strings.ToUpper("ImportWarning"):             MakePair(szW, 0x304),
		strings.ToUpper("PendingDeprecationWarning"): MakePair(szW, 0x305),
		strings.ToUpper("ResourceWarning"):           MakePair(szW, 0x306),
		strings.ToUpper("RuntimeWarning"):            MakePair(szW, 0x307),
		strings.ToUpper("SyntaxWarning"):             MakePair(szW, 0x308),
		strings.ToUpper("UnicodeWarning"):            MakePair(szW, 0x309),
		strings.ToUpper("UserWarning"):               MakePair(szW, 0x30A),
	}
)

type (
	// Clase base para todos los errores de Internet
	// Hereda de BaseError
	HTTPError struct {
		BaseError
	}

	// Clase base para los errores en databases
	// Hereda de BaseError
	DBError struct {
		BaseError
		Category string
	}

	// Clase base para los errores en ficheros
	// Hereda de BaseError
	FileIOError struct {
		BaseError
		Category string
	}

	customErr func(string, int, ...T) error
)

// Interface error
func (h *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: [%d] '%s'.", h.Code, h.Err)
}

// Interface error
func (db *DBError) Error() string {
	return fmt.Sprintf("DBError: In category '%s' [%d] '%s'.", db.Category, db.Code, db.Err)
}

// Interface error
func (fi *FileIOError) Error() string {
	return fmt.Sprintf("FileIOError: In category '%s' [%d] '%s'.", fi.Category, fi.Code, fi.Err)
}

// Devuelve un error HTTP
// @param {string}
// @param {...T} Se le pone por igualdad con las siguientes
// @return {error}
func NewHTTPError(code string, msg ...T) error {
	/* msg no usado*/
	pair := httpCodes[strings.ToUpper(code)]
	if pair == nil {
		return &HTTPError{BaseError{Code: -1, Err: "<HTTPERROR(" + code + ")>", Text: Literals().NullString}}
	}
	return &HTTPError{BaseError{Code: pair.T1.(int), Err: code, Text: pair.T2.(string)}}
}

// Devuelve un error DB
// @param {string}
// @param {...T}
// @return {error}
func NewDBError(code string, msg ...T) error {
	mlc := ArgOptional(Literals().NullString, msg).(string)
	pair := dbCodes[strings.ToUpper(code)]
	if pair == nil {
		return &DBError{BaseError{Code: -1, Err: "<DBERROR(" + code + ")>", Text: mlc}, Literals().NullString}
	}
	return &DBError{BaseError{Code: pair.T2.(int), Err: code, Text: mlc}, pair.T1.(string)}
}

// Devuelve un error FileIO
// @param {string}
// @param {...T}
// @return {error}
func NewFileIOError(code string, msg ...T) error {
	mlc := ArgOptional(Literals().NullString, msg).(string)
	pair := fileIOCodes[strings.ToUpper(code)]
	if pair == nil {
		return &FileIOError{BaseError{Code: -1, Err: "<FILEIOERROR(" + code + ")>" + code, Text: mlc}, Literals().NullString}
	}
	return &FileIOError{BaseError{Code: pair.T2.(int), Err: code, Text: mlc}, pair.T1.(string)}
}

// Emite un Warning
// @param {string}
// @param {...T}
// @return {error}
func NewWarning(code string, msg ...T) error {
	mlc := ArgOptional(Literals().NullString, msg).(string)
	pair := warningCodes[strings.ToUpper(code)]
	if pair == nil {
		return &Warning{BaseError{Code: -1, Err: "<WARNING(" + code + ")>", Text: mlc}}
	}
	return &Warning{BaseError{Code: pair.T2.(int), Err: code, Text: mlc}}
}

/** Emite un BaseError
* @param {string}
* @param {...T}
* @return {error}
 */
func NewBaseError(code string, msg ...T) error {
	mlc := ArgOptional(Literals().NullString, msg).(string)
	pair := baseErrorCodes[strings.ToUpper(code)]
	if pair == nil {
		return &BaseError{Code: -1, Err: "<BASEERROR(" + code + ")>", Text: strings.Trim(mlc+" ("+pair.T1.(string)+")", " ")}
	}
	return &BaseError{Code: pair.T2.(int), Err: code, Text: strings.Trim(mlc+" ("+pair.T1.(string)+")", " ")}
}

//!+
// Emite un Error Personalizado
// @param {string}
// @param {...T}
// @return {error}
// @@obsolete
func NewError(code string, nb int, msg ...T) error {
	mlc := ArgOptional(Literals().NullString, msg).(string)
	pair := new(Pair)
	pair.T1 = nb
	pair.T2 = mlc

	return &BaseError{Code: pair.T1.(int), Err: code, Text: pair.T2.(string)}
	//return fmt.Errorf("Error: '%s': '%d' : '%s'", code, pair.T1.(int), pair.T2.(string))
}

// Custom Error
var CustomError customErr = NewError

//!-
