package rex

// Next 用來將新的資料送入 Subject
func Next[A any](subject Subject[A]) func(a ...A) {
	next, _ := subject()

	return func(a ...A) {
		for _, a := range a {
			next <- ItemOf[A](a)
		}
	}
}

// NextItem 用來將新的 Item 資料送入 Subject
func NextItem[A any](subject Subject[A]) func(item ...Item[A]) {
	next, _ := subject()

	return func(item ...Item[A]) {
		for _, item := range item {
			next <- item
		}
	}
}

// NextChan 用來將新的 Channel 資料送入 Subject
func NextChan[A any](subject Subject[A]) func(ch <-chan A) {
	next, _ := subject()

	return func(ch <-chan A) {
		for a := range ch {
			next <- ItemOf[A](a)
		}
	}
}

// NextChanItem 用來將新的 Channel Item 資料送入 Subject
func NextChanItem[A any](subject Subject[A]) func(ch <-chan Item[A]) {
	next, _ := subject()

	return func(ch <-chan Item[A]) {
		for item := range ch {
			next <- item
		}
	}
}

// Subscribe 用來訂閱 Subject
func Subscribe[A any](subject Subject[A]) Iterable[A] {
	_, iterable := subject()

	return iterable
}
