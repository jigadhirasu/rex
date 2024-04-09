package rex

// Subject 是一個可以被訂閱的資料來源
func NewSubject[A any]() Subject[A] {
	next := make(chan Item[A])

	subscribes := []chan Item[A]{}
	closed := false

	go func() {
		defer func() {
			closed = true
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

		if closed {
			close(ch)
		}

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
