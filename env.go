/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/

package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type (
	/** @class */
	env           map[string]string
	sizeOfEnvFile int
	/** @class */
	sizeOfEnvs int
	/** @class */
	sep rune
	// struct
	separators struct {
		Separator string `json:"env.separator,omitempty"`
		Groups    bool   `json:"env.groups,omitempty"`
	}
)

var (
	//Environment container
	Env env = make(map[string]string)
	// Número de environments
	SizeOfEnvs sizeOfEnvs
	// Número de environments en el fichero
	SizeOfEnvFile sizeOfEnvFile
	// Separador
	separator sep
	seprtor   separators
	groups    bool
)

// TODO: Groups esta por hacer

func init() {
	seprtor = separators{}
	raw, _ := ioutil.ReadFile("./.envCfg.json")
	json.Unmarshal(raw, &seprtor)

	if seprtor.Separator == "colon" {
		separator = ':'
	} else {
		separator = '='
	}
	if seprtor.Groups {
		groups = true
	}
}

//interface Stringer
func (r *sep) String() string {
	return string(*r)
}

// Tamaño original del entorno
// @return {int}
func (s *sizeOfEnvs) SizeOfEnvsOriginal() int {
	return int(*s) - int(SizeOfEnvFile)
}

// Separador para el fichero '.envCfg.json'
// @return {string}
func (s *sizeOfEnvs) Separator() string {
	return string(separator)
}

// Grupos para el fichero '.envCfg.json'
// @return {bool}
func (s *sizeOfEnvs) Groups() bool {
	return groups
}

// Recupera las variables de entorno
func (e *env) Environ() {
	Env = make(map[string]string)
	s := os.Environ()
	SizeOfEnvs = sizeOfEnvs(len(s))

	for _, e := range s {
		pair := strings.SplitN(e, "=", 2)
		Env[pair[0]] = pair[1]
	}
}

// Recupera una variable
// @param {string}
// @return {string}
func (e *env) Get(s string) string {
	return Env[s]
}

// Coloca una variable
// @param {string}
// @param {string}
func (e *env) Set(k, s string) {
	// FIX: asegurarse que las envs son sólo para el proceso
	os.Setenv(k, s)
	Env[k] = s
	e.Environ()
}

// Carga un fichero '.env', ó el que sea, yml, etc
// @param {...string}
func (e *env) Load(file ...string) {
	if file == nil {
		file = make([]string, 1)
		file[0] = ".env"
	}
	readFile, err := os.Open(file[0])
	defer readFile.Close()
	if err != nil {
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	eq := fmt.Sprintf("%c", separator)

	var fileTextLines []string
	for fileScanner.Scan() {
		szAvg := fileScanner.Text()
		// Admite sólo si la línea es KEY=VALUE (ó KEY=).
		// Elimina las líneas en blanco ó los
		// comentarios de línea.
		if len(strings.SplitN(szAvg, eq, 2)) == 2 {
			fileTextLines = append(fileTextLines, szAvg)
		}
	}
	SizeOfEnvFile = sizeOfEnvFile(len(fileTextLines))
	for _, eachline := range fileTextLines {
		pair := strings.SplitN(eachline, eq, 2)
		// FIX: asegurarse que las envs son sólo para el proceso
		// FIX: que lo está ejecutando
		os.Setenv(pair[0], pair[1]) // las crea para el proceso
		Env[pair[0]] = pair[1]
	}

	e.Environ()
}
