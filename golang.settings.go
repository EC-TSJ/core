/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"maps"
	"os"
)

// *******************************************************************************
// ? @public @class Settings
// *******************************************************************************
type Settings map[string]string

var (
	// GlobalSettings
	GlobalSettings Settings = Settings{
		"STD":  "D:/x64/Go/src/std",
		"HOME": os.Getenv("HOME"),
		"core": "D:/x64/Go/src/core",
	}

	//!+
	__settings Settings = make(Settings)
)

// Ancho de GlobalSettings
// ? core.Method.Normal
// ! @return {int}
func (s Settings) Length() int {
	return len(GlobalSettings)
}

// Retorna un valor de GlobalSettings
// ? core.Method.Normal
// ! @return {int}
func (s Settings) Get(ut string) string {
	return GlobalSettings[ut]
}

/**
 * Crea un registro único, suma de los registros globales, locales, etc
 * Los valores de los registros de la derecha tienen preferencia sobre los valores de la izquierda, se
 * han de poner primero globales y luego locales.
 * <CODE>
 *?		golang.MakeSettings(GlobalSettings, LocalSettings)
 * </CODE>
 *! @param{...map[string]string} los registros a sumar
 */
func MakeSettings(m /*...Settings*/ Settings) {
	__settings := maps.Clone(GlobalSettings)
	maps.Copy(__settings, m)
}

/**
 * Obtiene un valor del registro de configuraciones
 *! @param{string}
 *! @return {any}
 */
func GetSetting(s string) string {
	return __settings[s]
}

//!-

// /**
//  * Obtiene el GlobalSettings
//  *! @return {Settings}
//  */
// func GSettings() Settings {
// 	return GlobalSettings
// }
