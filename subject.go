package rex

type Subject[A any] func() (chan<- Item[A], Iterable[A])

func (subject Subject[A]) Close() {
	next, _ := subject()
	close(next)
}

func NewSubject[A any]() Subject[A] {
	next := make(chan Item[A])

	subscribes := []chan Item[A]{}

	go func() {
		defer func() {
			for _, subscribe := range subscribes {
				close(subscribe)
			}
		}()

		for {
			item, ok := <-next
			if !ok {
				return
			}
			for _, subscribe := range subscribes {
				subscribe <- item
			}
		}
	}()

	subscribeFunc := func() <-chan Item[A] {
		source := make(chan Item[A])
		subscribes = append(subscribes, source)

		ch := make(chan Item[A])

		go func() {
			defer close(ch)

			for item := range source {
				ch <- item
			}
		}()

		return ch
	}

	return func() (chan<- Item[A], Iterable[A]) {
		return next, subscribeFunc
	}
}

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
