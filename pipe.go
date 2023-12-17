package rex

type PipeLine[A, B any] func(iterable Iterable[A]) func(ctx Context) Iterable[B]

func Pipe1[A, B any](p1 PipeLine[A, B]) PipeLine[A, B] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[B] {
		return func(ctx Context) Iterable[B] {
			return p1(iterable)(ctx)
		}
	}
}

func Pipe2[A, B, C any](p1 PipeLine[A, B], p2 PipeLine[B, C]) PipeLine[A, C] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[C] {
		return func(ctx Context) Iterable[C] {
			s1 := p1(iterable)(ctx)
			return p2(s1)(ctx)
		}
	}
}

func Pipe3[A, B, C, D any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D]) PipeLine[A, D] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[D] {
		return func(ctx Context) Iterable[D] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			return p3(s2)(ctx)
		}
	}
}

func Pipe4[A, B, C, D, E any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E]) PipeLine[A, E] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[E] {
		return func(ctx Context) Iterable[E] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			return p4(s3)(ctx)
		}
	}
}

func Pipe5[A, B, C, D, E, F any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F]) PipeLine[A, F] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[F] {
		return func(ctx Context) Iterable[F] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			s4 := p4(s3)(ctx)
			return p5(s4)(ctx)
		}
	}
}

func Pipe6[A, B, C, D, E, F, G any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G]) PipeLine[A, G] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[G] {
		return func(ctx Context) Iterable[G] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			s4 := p4(s3)(ctx)
			s5 := p5(s4)(ctx)
			return p6(s5)(ctx)
		}
	}
}

func Pipe7[A, B, C, D, E, F, G, H any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H]) PipeLine[A, H] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[H] {
		return func(ctx Context) Iterable[H] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			s4 := p4(s3)(ctx)
			s5 := p5(s4)(ctx)
			s6 := p6(s5)(ctx)
			return p7(s6)(ctx)
		}
	}
}

func Pipe8[A, B, C, D, E, F, G, H, I any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I]) PipeLine[A, I] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[I] {
		return func(ctx Context) Iterable[I] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			s4 := p4(s3)(ctx)
			s5 := p5(s4)(ctx)
			s6 := p6(s5)(ctx)
			s7 := p7(s6)(ctx)
			return p8(s7)(ctx)
		}
	}
}

func Pipe9[A, B, C, D, E, F, G, H, I, J any](p1 PipeLine[A, B], p2 PipeLine[B, C], p3 PipeLine[C, D], p4 PipeLine[D, E], p5 PipeLine[E, F], p6 PipeLine[F, G], p7 PipeLine[G, H], p8 PipeLine[H, I], p9 PipeLine[I, J]) PipeLine[A, J] {
	return func(iterable Iterable[A]) func(ctx Context) Iterable[J] {
		return func(ctx Context) Iterable[J] {
			s1 := p1(iterable)(ctx)
			s2 := p2(s1)(ctx)
			s3 := p3(s2)(ctx)
			s4 := p4(s3)(ctx)
			s5 := p5(s4)(ctx)
			s6 := p6(s5)(ctx)
			s7 := p7(s6)(ctx)
			s8 := p8(s7)(ctx)
			return p9(s8)(ctx)
		}
	}
}
