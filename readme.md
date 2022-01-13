# **Core**

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Core es una librería para manipular la datos básicos de la aplicación. Gestiona los tipos de datos, los errores y las variables de entorno.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales
---
--- 

Tiene los tipos siguientes:

 - ***BaseError*** `struct` 
 - ***Warning*** `struct` 
 - ***HTTPError*** `struct` 
 - ***DBError*** `struct` 
 - ***FileIOError*** `struct` 
 - ***Pair*** `struct`
 - ***NAFSS*** `struct`
 - ***CCC*** `struct`
 - ***CIF*** `struct`
 - ***NIE*** `struct`
 - ***NIF*** `struct`
 - ***CARD*** `struct`
 - ***IBAN*** `struct`

y otros tipos:

 - ***Item***  `interface{}`
 - ***Value***  `interface{}`
 - ***Key***   `interface{}`
 - ***Exception*** `interface{}`
 - ***TryCatch*** `struct`
 - ***While*** `struct`


Funciones: 

 - *`NewPair(Item, Item) *Pair`* 
 - *`MakePair(Item, Item) *Pair`*
 - *`NewHTTPError(string) error`* 
 - *`NewDBError(string) error`* 
 - *`NewFileIOError(string) error`* 
 - *`NewWarning(string) error`* 
 - *`NewBaseError(string) error`*
 - *`IIf(bool, interface(), interface())`*
 - *`If(bool, func() interface{}, func() interface{}) interface{}`*
 - *`While()`*
 - *`DoWhile()`*
 - *`TryCatch()`*
 - *`Throw()`*
 - *`ModifyString(*string, uintptr, int)`*
 - *`ModifySlice(unsafe.Pointer, uintptr, int)`*
 - *`ArgOptional(interface{}, ...interface{}) interface{}`*
   
   y argumentos por codificación

## Ejemplo
``` go

    msg = append(msg, nil)
    msg[0] = map[bool]interface{}{true: msg[0], false: "lo que sea"}[msg[0] != nil]
    msg = msg[:1]
    

```
 
 Variables: 

 - ***`Env`***           `map[string]string`
 - ***`SizeOfEnvs`***    `int`
 - ***`SizeOfEnvFile`*** `int`
 - ***`Code`***          `string`

objeto ***Env***, nos mapea los datos a un objeto map, con métodos:
 - *`Environ()`* 
 - *`Get(string) string `*
 - *`Set(string, string) `*
 - *`Load(...string) `*

objeto ***SizeOfEnvs***, con métodos:
 - *`SizeOfEnvsOriginal() int`*
 - *`Separator()`*
 - *`Groups()`*

objeto ***Code***, con métodos:
 - *`GetNIF(string) (bool, *NIF)`*
 - *`GetNIE(string) (bool, *NIE)`*
 - *`GetCCC(string) (bool, *CCC)`*
 - *`GetIBAN(string) (bool, *IBAN)`*
 - *`GetCARD(string) (bool, *CARD)`*
 - *`GetNAFSS(string [, bool]) (bool, *NAFSS)`*
 - *`GetCIF(string) (bool, *CIF)`*

y objeto ***CARD***, con método:
 - *`Card() string`*


## Ejemplos
```go

	var valu5 bool
	var pkop *core.NIE
	var pkip *core.NIF
	var sepp *core.CIF
	var seppA *core.NAFSS
	var seppB *core.CARD
	var seppC *core.IBAN
	var seppD *core.CCC

	valu5, seppD = core.Code.GetCCC("28797245373")
	valu5, seppD = core.Code.GetCCC("28797245317")
	fmt.Println(seppD, core.Code)
	valu5, seppC = core.Code.GetIBAN("ES3321003991000100134493" /*"14651234461234567890"*/)
	valu5, seppC = core.Code.GetIBAN("ES3321003991480100134493" /*"14651234461234567890"*/)
	fmt.Println(seppC, core.Code)
	valu5, seppB = core.Code.GetCARD("5489133149518807")
	valu5, seppB = core.Code.GetCARD("5489133149518803")
	fmt.Println(seppB, core.Code, seppB.Card())
	valu5, seppA = core.Code.GetNAFSS("280383841421")
	fmt.Println(seppA, core.Code)
	valu5, sepp = core.Code.GetCIF("A58818501")
	fmt.Println(sepp, core.Code)
	valu5, sepp = core.Code.GetCIF("A78145398")
	fmt.Println(sepp, core.Code)
	valu5, sepp = core.Code.GetCIF("B83524868")
	fmt.Println(sepp, core.Code)
	valu5, sepp = core.Code.GetCIF("W83524868")
	fmt.Println(sepp, core.Code)
	valu5, pkip = core.Code.GetNIF("50631377")
	fmt.Println(pkip, core.Code)
	valu5, pkip = core.Code.GetNIF("50031377Y")
	fmt.Println(pkip, core.Code)
	valu5, pkip = core.Code.GetNIF("50025660B")
	fmt.Println(pkip, core.Code)
	valu5, pkip = core.Code.GetNIF("50022412Q")
	fmt.Println(pkip, core.Code)
	valu5, pkop = core.Code.GetNIE("X1595989")
	fmt.Println(pkop, core.Code)
	valu5, pkop = core.Code.GetNIE("X1771291")
	fmt.Println(pkop, core.Code)



	core.Env.Environ()
	swkw := core.NewBaseError("ArithmeticError")
	swkw1 := core.NewFileIOError("BrokenPipeError", "** Loadfactor por debajo del límite, aumentando tamaño **")
	swkw12 := core.NewBaseError("ArithmeticError", "** Loadfactor por debajo del límite, aumentando tamaño **")
	swkw13 := core.NewBaseError("ArithmeticError24", "** Loadfactor por debajo del límite, aumentando tamaño **")
	per := core.Env.Get("INCLUDE")
	core.Env.Load(".env")
	asir := core.SizeOfEnvFile
	asirA := core.SizeOfEnvs
	asirB := core.SizeOfEnvs.SizeOfEnvsOriginal()
	asirC := core.SizeOfEnvs.Separator()
	jjj := core.Env["MARVI"]

	meda := core.IIf(true, Data(), -15)
	core.If(true, func() error {
		fmt.Println("++++")
		return nil
	}, func() error {
		fmt.Println("ttttt")
		return nil
	})
	core.If(false, func() error {
		fmt.Println("++++")
		return nil
	}, func() error {
		fmt.Println("ttttt")
		return nil
	})


```
## Notas





<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
