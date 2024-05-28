package core

type (
	ClassMethods byte
	ClassType    byte
	MethodType   byte
)

const (
	_Class_     ClassMethods = 0
	_Method_    ClassMethods = 1
	_Abstract_  ClassType    = 0b0000
	_Derived_   ClassType    = 0b0001
	_Normal_    ClassType    = 1 << 1
	_Sealed_    ClassType    = 1 << 2
	_Protected_ ClassType    = 1 << 3
	_Virtual_   MethodType   = 1 << 0
	_Override_  MethodType   = 1 << 1
	_Final_     MethodType   = 1 << 2
)

/**
 *? @public @enum ClassMethod
 */
var ClassMethod = &struct {
	Class, Method ClassMethods
}{
	Class:  _Class_,
	Method: _Method_,
}

/**
 *? @public @enum Class
 */
var Class = &struct {
	Abstract, Derived, Normal, Sealed, Protected ClassType
}{
	Abstract:  _Abstract_,
	Derived:   _Derived_,
	Normal:    _Normal_,
	Sealed:    _Sealed_,
	Protected: _Protected_,
}

/**
 *? @public @enum Method
 */
var Method = &struct {
	Virtual, Override, Final MethodType
}{
	Virtual:  _Virtual_,
	Override: _Override_,
	Final:    _Final_,
}

func (v ClassType) String() string {
	return [...]string{
		"Abstract Class", //0
		"Derived Class",  //1
		"Normal Class",   //2
		"",
		"Sealed Class", //4
		"",
		"",
		"",
		"Protected Class", //8
	}[int(v)]
}

func (v MethodType) String() string {
	return [...]string{
		"Normal Method",   //0
		"Virtual Method",  //1
		"Override Method", //2
		"",
		"Final Method", //4
	}[int(v)]
}

func (v ClassMethods) String() string {
	return [...]string{
		"Class Object",  //0
		"Method Object", //1
	}[int(v)]
}
