/**
 * Funcion que realiza una llamada y devuelve bool
 */
package core

type PredicateV0 func() bool

type PredicateV1[T1 any] func(T1) bool
type PredicateV2[T1 any, T2 any] func(T1, T2) bool
type PredicateV3[T1 any, T2 any, T3 any] func(T1, T2, T3) bool
type PredicateV4[T1 any, T2 any, T3 any, T4 any] func(T1, T2, T3, T4) bool
type PredicateV5[T1 any, T2 any, T3 any, T4 any, T5 any] func(T1, T2, T3, T4, T5) bool

type PredicateV1X[T1 any] func(...T1) bool
type PredicateV2X[T1 any, T2 any] func(T1, ...T2) bool
type PredicateV3X[T1 any, T2 any, T3 any] func(T1, T2, ...T3) bool
type PredicateV4X[T1 any, T2 any, T3 any, T4 any] func(T1, T2, T3, ...T4) bool
type PredicateV5X[T1 any, T2 any, T3 any, T4 any, T5 any] func(T1, T2, T3, T4, ...T5) bool
