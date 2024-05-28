/**
 * Funcion que realiza una llamada y devuelve el ultimo tipo definido como TResult
 */

package core

type FuncV0[TResult any] func() TResult

type FuncV1[T1 any, TResult any] func(T1) TResult
type FuncV2[T1 any, T2 any, TResult any] func(T1, T2) TResult
type FuncV3[T1 any, T2 any, T3 any, TResult any] func(T1, T2, T3) TResult
type FuncV4[T1 any, T2 any, T3 any, T4 any, TResult any] func(T1, T2, T3, T4) TResult
type FuncV5[T1 any, T2 any, T3 any, T4 any, T5 any, TResult any] func(T1, T2, T3, T4, T5) TResult

type FuncV1X[T1 any, TResult any] func(...T1) TResult
type FuncV2X[T1 any, T2 any, TResult any] func(T1, ...T2) TResult
type FuncV3X[T1 any, T2 any, T3 any, TResult any] func(T1, T2, ...T3) TResult
type FuncV4X[T1 any, T2 any, T3 any, T4 any, TResult any] func(T1, T2, T3, ...T4) TResult
type FuncV5X[T1 any, T2 any, T3 any, T4 any, T5 any, TResult any] func(T1, T2, T3, T4, ...T5) TResult
