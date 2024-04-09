package rex

// Reader 用來定義一個資料流的處理流程
type Reader[A any] func(ctx Context) Iterable[A]

// PipeLine 接收一個資料流後，產生新的資料流
type PipeLine[A, B any] func(iterable Iterable[A]) Reader[B]

// Pipe1 用來將一個 PipeLine 與一個 Iterable 組合成一個 Reader
func Pipe1[A, B any](p1 PipeLine[A, B]) PipeLine[A, B] {
	return func(iterable Iterable[A]) Reader[B] {
		return func(ctx Context) Iterable[B] {
			return p1(iterable)(ctx)
		}
	}
}

// Pipe2 用來將兩個 PipeLine 組合成一個 PipeLine
func Pipe2[A, B, C any](p1 PipeLine[A, B], p2 PipeLine[B, C]) PipeLine[A, C] {
	return func(iterable Iterable[A]) Reader[C] {
		return func(ctx Context) Iterable[C] {
			return p2(Pipe1[A, B](p1)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe3 用來將三個 PipeLine 組合成一個 PipeLine
func Pipe3[A, B, C, D any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D]) PipeLine[A, D] {
	return func(iterable Iterable[A]) Reader[D] {
		return func(ctx Context) Iterable[D] {
			return p3(Pipe2[A, B, C](p1, p2)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe4 用來將四個 PipeLine 組合成一個 PipeLine
func Pipe4[A, B, C, D, E any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E]) PipeLine[A, E] {
	return func(iterable Iterable[A]) Reader[E] {
		return func(ctx Context) Iterable[E] {
			return p4(Pipe3[A, B, C, D](p1, p2, p3)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe5 用來將五個 PipeLine 組合成一個 PipeLine
func Pipe5[A, B, C, D, E, F any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F]) PipeLine[A, F] {
	return func(iterable Iterable[A]) Reader[F] {
		return func(ctx Context) Iterable[F] {
			return p5(Pipe4[A, B, C, D, E](p1, p2, p3, p4)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe6 用來將六個 PipeLine 組合成一個 PipeLine
func Pipe6[A, B, C, D, E, F, G any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G]) PipeLine[A, G] {
	return func(iterable Iterable[A]) Reader[G] {
		return func(ctx Context) Iterable[G] {
			return p6(Pipe5[A, B, C, D, E, F](p1, p2, p3, p4, p5)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe7 用來將七個 PipeLine 組合成一個 PipeLine
func Pipe7[A, B, C, D, E, F, G, H any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H]) PipeLine[A, H] {
	return func(iterable Iterable[A]) Reader[H] {
		return func(ctx Context) Iterable[H] {
			return p7(Pipe6[A, B, C, D, E, F, G](p1, p2, p3, p4, p5, p6)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe8 用來將八個 PipeLine 組合成一個 PipeLine
func Pipe8[A, B, C, D, E, F, G, H, I any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I]) PipeLine[A, I] {
	return func(iterable Iterable[A]) Reader[I] {
		return func(ctx Context) Iterable[I] {
			return p8(Pipe7[A, B, C, D, E, F, G, H](p1, p2, p3, p4, p5, p6, p7)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe9 用來將九個 PipeLine 組合成一個 PipeLine
func Pipe9[A, B, C, D, E, F, G, H, I, J any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J]) PipeLine[A, J] {
	return func(iterable Iterable[A]) Reader[J] {
		return func(ctx Context) Iterable[J] {
			return p9(Pipe8[A, B, C, D, E, F, G, H, I](p1, p2, p3, p4, p5, p6, p7, p8)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe10 用來將十個 PipeLine 組合成一個 PipeLine
func Pipe10[A, B, C, D, E, F, G, H, I, J, K any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K]) PipeLine[A, K] {
	return func(iterable Iterable[A]) Reader[K] {
		return func(ctx Context) Iterable[K] {
			return p10(Pipe9[A, B, C, D, E, F, G, H, I, J](p1, p2, p3, p4, p5, p6, p7, p8, p9)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe11 用來將十一個 PipeLine 組合成一個 PipeLine
func Pipe11[A, B, C, D, E, F, G, H, I, J, K, L any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L]) PipeLine[A, L] {
	return func(iterable Iterable[A]) Reader[L] {
		return func(ctx Context) Iterable[L] {
			return p11(Pipe10[A, B, C, D, E, F, G, H, I, J, K](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe12 用來將十二個 PipeLine 組合成一個 PipeLine
func Pipe12[A, B, C, D, E, F, G, H, I, J, K, L, M any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M]) PipeLine[A, M] {
	return func(iterable Iterable[A]) Reader[M] {
		return func(ctx Context) Iterable[M] {
			return p12(Pipe11[A, B, C, D, E, F, G, H, I, J, K, L](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe13 用來將十三個 PipeLine 組合成一個 PipeLine
func Pipe13[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N]) PipeLine[A, N] {
	return func(iterable Iterable[A]) Reader[N] {
		return func(ctx Context) Iterable[N] {
			return p13(Pipe12[A, B, C, D, E, F, G, H, I, J, K, L, M](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe14 用來將十四個 PipeLine 組合成一個 PipeLine
func Pipe14[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O]) PipeLine[A, O] {
	return func(iterable Iterable[A]) Reader[O] {
		return func(ctx Context) Iterable[O] {
			return p14(Pipe13[A, B, C, D, E, F, G, H, I, J, K, L, M, N](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe15 用來將十五個 PipeLine 組合成一個 PipeLine
func Pipe15[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O], p15 PipeLine[O, P]) PipeLine[A, P] {
	return func(iterable Iterable[A]) Reader[P] {
		return func(ctx Context) Iterable[P] {
			return p15(Pipe14[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe16 用來將十六個 PipeLine 組合成一個 PipeLine
func Pipe16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O], p15 PipeLine[O, P], p16 PipeLine[P, Q]) PipeLine[A, Q] {
	return func(iterable Iterable[A]) Reader[Q] {
		return func(ctx Context) Iterable[Q] {
			return p16(Pipe15[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe17 用來將十七個 PipeLine 組合成一個 PipeLine
func Pipe17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O], p15 PipeLine[O, P], p16 PipeLine[P, Q], p17 PipeLine[Q, R]) PipeLine[A, R] {
	return func(iterable Iterable[A]) Reader[R] {
		return func(ctx Context) Iterable[R] {
			return p17(Pipe16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe18 用來將十八個 PipeLine 組合成一個 PipeLine
func Pipe18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O], p15 PipeLine[O, P], p16 PipeLine[P, Q], p17 PipeLine[Q, R], p18 PipeLine[R, S]) PipeLine[A, S] {
	return func(iterable Iterable[A]) Reader[S] {
		return func(ctx Context) Iterable[S] {
			return p18(Pipe17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17)(iterable)(ctx))(ctx)
		}
	}
}

// Pipe19 用來將十九個 PipeLine 組合成一個 PipeLine
func Pipe19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J], p10 PipeLine[J, K], p11 PipeLine[K, L], p12 PipeLine[L, M], p13 PipeLine[M, N], p14 PipeLine[N, O], p15 PipeLine[O, P], p16 PipeLine[P, Q], p17 PipeLine[Q, R], p18 PipeLine[R, S], p19 PipeLine[S, T]) PipeLine[A, T] {
	return func(iterable Iterable[A]) Reader[T] {
		return func(ctx Context) Iterable[T] {
			return p19(Pipe18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R](p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18)(iterable)(ctx))(ctx)
		}
	}
}
