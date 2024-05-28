/**
 * Funcion que realiza una llamada y no devuelve nada
 */

package core

type ActionV0 func()

type ActionV1[T1 any] func(T1)
type ActionV2[T1 any, T2 any] func(T1, T2)
type ActionV3[T1 any, T2 any, T3 any] func(T1, T2, T3)
type ActionV4[T1 any, T2 any, T3 any, T4 any] func(T1, T2, T3, T4)
type ActionV5[T1 any, T2 any, T3 any, T4 any, T5 any] func(T1, T2, T3, T4, T5)

type ActionV1X[T1 any] func(...T1)
type ActionV2X[T1 any, T2 any] func(T1, ...T2)
type ActionV3X[T1 any, T2 any, T3 any] func(T1, T2, ...T3)
type ActionV4X[T1 any, T2 any, T3 any, T4 any] func(T1, T2, T3, ...T4)
type ActionV5X[T1 any, T2 any, T3 any, T4 any, T5 any] func(T1, T2, T3, T4, ...T5)
