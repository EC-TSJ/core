/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"fmt"
	"math"
	"strconv"
)

var (
	_MODULEBASE_NIF   int = 23
	_MODULEBASE_NAFSS int = 97
	_WIDTH            int = 24

	letras = []rune{'T', 'R', 'W', 'A', 'G', 'M', 'Y', 'F', 'P', 'D', 'X', 'B', 'N', 'J', 'Z', 'S', 'Q', 'V', 'H', 'L', 'C', 'K', 'E'}
	pesos  = []int{6, 3, 7, 9, 10, 5, 8, 4, 2, 1}
	// Codes container
	Code code
)

type (
	code string

	_err struct {
		Msg error
	}
	// Tipo CIF
	CIF struct {
		_err
		Tipo   string
		Numero string
		DC     string
	}
	// Tipo NIE
	NIE struct {
		_err
		Tipo   string
		Numero string
		DC     string
	}
	// Tipo NIF
	NIF struct {
		_err
		Numero string
		DC     string
	}
	// Tipo CCC
	CCC struct {
		_err
		Provincia string
		Numero    string
		DC        string
	}
	// Tipo NAFSS
	NAFSS struct {
		_err
		Provincia string
		Numero    string
		DC        string
	}
	// Tipo CARD
	CARD struct {
		_err
		Code  string
		Code1 string
		Code2 string
		Code3 string
		Code4 string
	}
	// Tipo IBAN
	IBAN struct {
		_err
		Iban    string
		Entidad string
		Oficina string
		DC      string
		CtaCte  string
	}

	// Interface ICif
	ICif interface {
		Cif() string
	}

	// Interface ICard
	ICard interface {
		Card() string
	}
)

//!+
/**
 * @brief GetCIF Obtiene un CIF de una empresa
 * @param {string}
 * @return {bool, *CIF}
 */
func (*code) GetCIF(sCIF string) (bool, *CIF) {
	var sumP, sumI int = 0, 0
	var flag bool = false
	mapy := map[byte]byte{
		'0': 'J',
		'1': 'A',
		'2': 'B',
		'3': 'C',
		'4': 'D',
		'5': 'E',
		'6': 'F',
		'7': 'G',
		'8': 'H',
		'9': 'I',
	}

	var sLetra byte = sCIF[0]
	var sMid string = sCIF[1:8]
	// generan un caracter final y no un digito: P Q R S W
	if sLetra == 'P' || sLetra == 'Q' || sLetra == 'R' || sLetra == 'S' || sLetra == 'W' || sMid[:2] == "00" {
		flag = true
	}

	// suma los valores pares e impares de la cifra
	var par bool = true // false = impar; true = par; impar por defecto
	for _, v := range sMid {
		v -= 48
		if !par {
			// suma los pares
			sumP += int(v)
		} else {
			// suma los impares
			v *= 2
			if v >= 10 {
				// suma los dos caracteres: 16 = 1 + 6 = 7
				v = rune(strconv.Itoa(int(v))[0]-48) + rune(strconv.Itoa(int(v))[1]-48)
			}
			sumI += int(v)
		}
		// hace par/impar
		par = !par
	}
	sum := sumI + sumP // hace el total de las sumas

	var caracter byte
	if sum >= 10 {
		// se queda con el segundo caracter, el de las unidades
		caracter = strconv.Itoa(sum)[1]
	} else {
		caracter = byte(sum) + 48
	}

	if caracter != 48 {
		var value byte = caracter - 48
		// hace la operación para el calculo del cif
		caracter = 10 - value
		value = caracter + 48
		caracter = value
	} else {
		caracter = byte(48)
	}

	var sDesc byte
	if flag { // añade el dígito/letra
		sDesc = mapy[caracter]
	} else {
		sDesc = caracter
	}

	Code = code(string(sLetra) + sMid + string(sDesc))
	return /*sCIF == string(Code)*/ true, &CIF{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Tipo:   string(Code[0]),
		Numero: string(Code[1 : len(Code)-1]),
		DC:     string(Code[len(Code)-1])}
}

/** Interface ICif */
func (cif *CIF) Cif() string {
	switch cif.Tipo {
	case "A":
		return fmt.Sprintf("%s", "Sociedad Anónima")
	case "B":
		return fmt.Sprintf("%s", "Sociedad Responsabilidad Limitada")
	case "C":
		return fmt.Sprintf("%s", "Sociedad Colectiva")
	case "D":
		return fmt.Sprintf("%s", "Sociedad Comanditaria")
	case "E":
		return fmt.Sprintf("%s", "Comunidad de Bienes")
	case "F":
		return fmt.Sprintf("%s", "Sociedad Cooperativa")
	case "G":
		return fmt.Sprintf("%s", "Asociación")
	case "H":
		return fmt.Sprintf("%s", "Comunidad de Propietarios")
	case "J":
		return fmt.Sprintf("%s", "Sociedad civil")
	case "N":
		return fmt.Sprintf("%s", "Sociedad extranjera")
	case "P":
		return fmt.Sprintf("%s", "Corporación local")
	case "Q":
		return fmt.Sprintf("%s", "Organismo Público")
	case "R":
		return fmt.Sprintf("%s", "Congregación/Institución Religiosa")
	case "S":
		return fmt.Sprintf("%s", "Órgano de la Administración y de CCAA") //Administración General del Estado y de CCAA
	case "U":
		return fmt.Sprintf("%s", "Union Temporal de Empresa")
	case "V":
		return fmt.Sprintf("%s", "Otro tipo no definido")
	case "W":
		return fmt.Sprintf("%s", "Establecimiento Permanente") // de Entidad no Residente en España
	default:
		return fmt.Sprintf("")
	}
}

/** Interface Stringer */
func (cif *CIF) String() string {
	if w, ok := T(cif).(ICif); ok {
		return fmt.Sprintf("%s (%s/%s%s)", w.Cif(), cif.Tipo, cif.Numero, cif.DC)
	}
	return fmt.Sprintf("%s/%s%s", cif.Tipo, cif.Numero, cif.DC)
}

//!-

//!+
/**
 * @brief GetNIE Devuelve el NIE correspondiente a un DNI
 * @param {string}
 * @return {bool, *NIE}
 */
func (nie *code) GetNIE(sNIE string) (bool, *NIE) {
	x := sNIE[0]
	sNIE = sNIE[1:]

	if x != 'X' && x != 'Y' && x != 'Z' {
		return false, &NIE{_err: _err{Msg: NewBaseError("ParameterError", "** Tipo de NIF _erróneo **")}, Tipo: "", Numero: "", DC: ""}
	}
	/* Code = */ nie.GetNIF(sNIE) // deja el dato en Code

	Code = code(x) + Code
	return true, &NIE{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Tipo:   string(Code[0]),
		Numero: string(Code[1 : len(Code)-1]),
		DC:     string(Code[len(Code)-1])}
}

/** Interface Stringer */
func (nie *NIE) String() string {
	return fmt.Sprintf("%s-%s-%s", nie.Tipo, nie.Numero, nie.DC)
}

//!-

//!+
/**
 * @brief GetNIF Devuelve el NIF correspondiente a un DNI
 * @param {string}
 * @return {bool, *NIF}
 */
func (*code) GetNIF(sDNI string) (bool, *NIF) {
	tmp := sDNI
	Code = ""

	if tmp[len(tmp)-1] > 57 { // si va la letra puesta
		tmp = tmp[0 : len(tmp)-1]
	}
	iDNI, _ := strconv.Atoi(tmp)
	// averigua la letra
	if iDNI < 100000000 && iDNI > _MODULEBASE_NIF {
		// obtiene el desplazamiento de la letra correspondiente,  dentro de la matriz letras[]
		var diff int = iDNI - int(math.Floor(float64(iDNI/_MODULEBASE_NIF)))*_MODULEBASE_NIF
		// Añade la letra
		tmp = tmp + string(letras[diff])
	}

	Code = code(tmp)
	return /*string(Code) == sDNI*/ true, &NIF{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Numero: string(Code[:len(Code)-1]),
		DC:     string(Code[len(Code)-1])}
}

/** Interface Stringer */
func (nif *NIF) String() string {
	return fmt.Sprintf("%s-%s", nif.Numero, nif.DC)
}

//!-

//!+
/**
 * @brief GetCCC Devuelve el CCC de la seg. social
 * @param {string}
 * @return {bool, *CCC}
 */
func (ccc *code) GetCCC(sCCC string) (bool, *CCC) {
	st := new(CCC)
	_, sta := ccc.GetNAFSS(sCCC, false)
	st.Provincia = sta.Provincia
	st.Numero = sta.Numero[0 : len(sta.Numero)-1]
	st.DC = sta.DC

	return sCCC[9:] == st.DC, &CCC{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Provincia: string(Code[:2]),
		Numero:    string(Code[2:9]),
		DC:        string(Code[len(Code)-2:])}
}

// Interface Stringer
func (ccc *CCC) String() string {
	return fmt.Sprintf("%s%s-%s", ccc.Provincia, ccc.Numero, ccc.DC)
}

//!-

//!+
/**
 * @brief GetNAFSS Devuelve el NAFSS correspondiente de la Seg. Social
 * @param {string}
 * @param {...bool} fMode true -> trabajador; false -> Empresa
 * @return {bool, *NAFSS}
 */
func (*code) GetNAFSS(sNAFSS string, fMode ...bool) (bool, *NAFSS) {
	if len(fMode) != 1 { // make fMode optional
		fMode = make([]bool, 1)
		fMode[0] = true
	}
	//program efective
	sA := sNAFSS[0:2]
	var value int = 9
	if fMode[0] {
		value = 10
	}
	sB := sNAFSS[2:value]
	sC := sB
	if sC[0] == '0' { //elimina el 0 de la tercera posición
		sC = sC[1:]
	}

	sNAFSS = sA + sC
	// Valor/97
	result, _ := strconv.ParseFloat(sNAFSS, 64)
	result = result / float64(_MODULEBASE_NAFSS)
	// Result = Result - (parte entera de Result)
	result -= math.Floor(result)
	// Result = Result * 100
	result *= 100

	// corrección de número y colocación del result en szBufferTmp
	if result <= 65 && result >= 33 {
		result -= 1
	}
	if result <= 98 && result >= 66 {
		result -= 2
	}

	// quita la parte decimal
	intpart, partfract := math.Modf(result)
	result -= partfract

	// convierte los digitos de control
	// y lo añade al código
	sD := strconv.Itoa(int(intpart))
	sNAFSS = sA + sB + sD
	Code = code(sNAFSS)

	return sNAFSS == string(Code), &NAFSS{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Provincia: string(Code[:2]),
		Numero:    string(Code[2:10]),
		DC:        string(Code[len(Code)-2:])}
}

// Interface Stringer
func (nafss *NAFSS) String() string {
	return fmt.Sprintf("%s%s-%s", nafss.Provincia, nafss.Numero, nafss.DC)
}

//!-

//!+
/**
 * @brief GetCARD Comprueba si el numero de tarjeta es correcto
 * @param {string}
 * @return {bool, *CARD}
 */
func (*code) GetCARD(sCardNumber string) (bool, *CARD) {
	bChar := sCardNumber[0] //primer caracter de la tarjeta (marca de tarjeta)

	/* controla que sea un número de un ancho correcto (16 digitos) y
	   del tipo correcto */
	if len(sCardNumber) != 16 || (bChar != '3' && bChar != '4' && bChar != '5' && bChar != '6') {
		Code = code(sCardNumber)
		return false, &CARD{_err: _err{Msg: NewBaseError("ParameterError", "** Error en el número de Tarjeta **")},
			Code: string(Code[0]), Code1: string(Code[1:4]), Code2: string(Code[4:8]), Code3: string(Code[8:12]), Code4: string(Code[12:16])}
	}

	/* recorre todos los elementos de la cadena para ver si son
	   pares o impares y hacer la suma */
	var sumI, sumP int = 0, 0
	for a := range sCardNumber {
		_, remainder := _ldiv(int64(a+1), 2) //son pares ó impares
		if remainder == 0 {
			//pares
			sumP += (int(sCardNumber[a]) - 48)
		} else {
			// impares
			calc := (int(sCardNumber[a]) - 48) * 2
			if calc > 9 {
				calc -= 9
			}
			sumI += calc
		}
	}
	total := sumI + sumP

	// comprueba que es multiplo de 10 y menor/igual que 150
	_, remainder := _ldiv(int64(total), 10)
	Code = code(sCardNumber)
	if (remainder == 0) && (total <= 150) {
		return true, &CARD{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
			Code:  string(Code[0]),
			Code1: string(Code[1:4]),
			Code2: string(Code[4:8]),
			Code3: string(Code[8:12]),
			Code4: string(Code[12:16])}
	} else {
		return false, &CARD{_err: _err{Msg: NewBaseError("ParameterError", "** Error en el número de Tarjeta **")},
			Code:  string(Code[0]),
			Code1: string(Code[1:4]),
			Code2: string(Code[4:8]),
			Code3: string(Code[8:12]),
			Code4: string(Code[12:16])}
	}
}

func _ldiv(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator

	return
}

/** Interface Stringer */
func (card *CARD) String() string {
	// Ejecuta interface ICard
	if w, ok := T(card).(ICard); ok {
		return fmt.Sprintf("%s(%s%s-%s-%s-%s)", w.Card(), card.Code, card.Code1, card.Code2, card.Code3, card.Code4)
	}
	return fmt.Sprintf("%s%s-%s-%s-%s", card.Code, card.Code1, card.Code2, card.Code3, card.Code4)
}

/** Interface ICard */
func (card *CARD) Card() string {
	switch card.Code {
	case "3":
		return fmt.Sprintf("%s", "American Express")
	case "4":
		return fmt.Sprintf("%s", "Visa")
	case "5":
		return fmt.Sprintf("%s", "Master Card")
	case "6":
		return fmt.Sprintf("%s", "Discover")
	default:
		return ""
	}
}

//!-

//!+
/**
 * helper from GetIBAN
 * @brief GetCtaCte Obtiene un numero de Cuenta Corriente correcto
 * @param {string}
 * @return {string}
 */
func _getCtaCte(szCCC string) string {
	if len(szCCC) != 20 {
		return "err:Tamaño"
	}

	szBuffer := ""
	// 1ª parte
	digit := _digitCtaCte("00" + szCCC[0:8])
	if digit == -1 {
		return "err:Entidad/Oficina"
	} else {
		szBuffer = szCCC[:8] + string(rune(digit+48))
	}

	//segunda parte (Cuenta)
	digit = _digitCtaCte(szCCC[10:])
	if digit == -1 {
		return "err:Cta/Cte"
	} else {
		szBuffer = szBuffer[:9] + string(rune(digit+48))
	}
	szBuffer += szCCC[10:]

	return szBuffer
}

// helper from _getCtaCte
func _digitCtaCte(sCadena string) int {
	sum := 0
	lena := len(sCadena)

	// calcula el peso
	// a) suma de todos los productos
	for off := range sCadena {
		bChar := sCadena[lena-off-1]
		if bChar < 48 || bChar > 57 { // no es un numero
			return -1
		} else {
			sum = sum + (int(bChar-48) * pesos[off])
		}
	}

	// b) en base al modulo 11
	_, remainder := _ldiv(int64(sum), 11)
	sum = int(11 - remainder)

	//c)correccion
	switch sum {
	case 10:
		return 1
	case 11:
		return 0
	default:
		return sum
	}
}

/**
 * @brief GetIBAN Calcula el IBAN
 * @param {string}
 * @return {bool, *IBAN}
 */
func (iban *code) GetIBAN(sCuenta string) (bool, *IBAN) {
	if sCuenta[:2] == "ES" {
		sCuenta = sCuenta[4:]
	}

	stmpA := sCuenta
	sCuenta = _getCtaCte(sCuenta)
	stmpB := sCuenta
	sCuenta += "142800"

	// Hace el calculo del número en dos pasos
	result, _ := strconv.ParseInt(sCuenta[0:15], 10, 64)
	two := result % 97
	//---------
	var ll string = strconv.Itoa(int(two))
	sCuenta = ll + string(sCuenta[15:])
	//---------
	result, _ = strconv.ParseInt(sCuenta, 10, 64)
	two = result % 97
	// final del segundo paso -> cálculo del código
	var res int64 = 98 - two
	//Cálculo hecho

	var s string = strconv.Itoa(int(res))
	if len(s) == 1 {
		s = "0" + s
	}

	var sf string = "ES" + s + stmpB
	Code = code(sf)

	return sf[4:] == stmpA, &IBAN{_err: _err{Msg: NewBaseError("Correct", "OperationValided")},
		Iban:    string(Code[0:4]),
		Entidad: string(Code[4:8]),
		Oficina: string(Code[8:12]),
		DC:      string(Code[12:14]),
		CtaCte:  string(Code[14:24])}
}

/** Interface Stringer */
func (iban *IBAN) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", iban.Iban, iban.Entidad, iban.Oficina, iban.DC, iban.CtaCte)
}

//!-
