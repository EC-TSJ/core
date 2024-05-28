/**
 * Funcion que realiza una llamada y devuelve dos ultimos tipos definidos como TResult y TResultx
 */

package core

type RFuncV0[TResult any, TResultx any] func() (TResult, TResultx)

type RFuncV1[T1 any, TResult any, TResultx any] func(T1) (TResult, TResultx)
type RFuncV2[T1 any, T2 any, TResult any, TResultx any] func(T1, T2) (TResult, TResultx)
type RFuncV3[T1 any, T2 any, T3 any, TResult any, TResultx any] func(T1, T2, T3) (TResult, TResultx)
type RFuncV4[T1 any, T2 any, T3 any, T4 any, TResult any, TResultx any] func(T1, T2, T3, T4) (TResult, TResultx)
type RFuncV5[T1 any, T2 any, T3 any, T4 any, T5 any, TResult any, TResultx any] func(T1, T2, T3, T4, T5) (TResult, TResultx)

type RFuncV1X[T1 any, TResult any, TResultx any] func(...T1) (TResult, TResultx)
type RFuncV2X[T1 any, T2 any, TResult any, TResultx any] func(T1, ...T2) (TResult, TResultx)
type RFuncV3X[T1 any, T2 any, T3 any, TResult any, TResultx any] func(T1, T2, ...T3) (TResult, TResultx)
type RFuncV4X[T1 any, T2 any, T3 any, T4 any, TResult any, TResultx any] func(T1, T2, T3, ...T4) (TResult, TResultx)
type RFuncV5X[T1 any, T2 any, T3 any, T4 any, T5 any, TResult any, TResultx any] func(T1, T2, T3, T4, ...T5) (TResult, TResultx)
