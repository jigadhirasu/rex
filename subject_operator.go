package rex

func Next[A any](subject Subject[A]) func(a ...A) {
	next, _ := subject()

	return func(a ...A) {
		for _, a := range a {
			next <- ItemOf[A](a)
		}
	}
}

func NextItem[A any](subject Subject[A]) func(item ...Item[A]) {
	next, _ := subject()

	return func(item ...Item[A]) {
		for _, item := range item {
			next <- item
		}
	}
}

func NextChan[A any](subject Subject[A]) func(ch <-chan A) {
	next, _ := subject()

	return func(ch <-chan A) {
		for a := range ch {
			next <- ItemOf[A](a)
		}
	}
}

func NextChanItem[A any](subject Subject[A]) func(ch <-chan Item[A]) {
	next, _ := subject()

	return func(ch <-chan Item[A]) {
		for item := range ch {
			next <- item
		}
	}
}

func Subscribe[A any](subject Subject[A]) Iterable[A] {
	_, iterable := subject()

	return iterable
}
