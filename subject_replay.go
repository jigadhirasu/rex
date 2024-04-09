package rex

import (
	"sync"
)

// NewSubjectReplay 用來建立一個帶有回放功能的 Subject
func NewSubjectReplay[A any](replayCount int) Subject[A] {
	next := make(chan Item[A])

	subscribes := []chan Item[A]{}

	replay := []A{}
	replayMutex := new(sync.RWMutex)
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

			a, err := item()
			if err != nil {
				return
			}

			replayMutex.Lock()
			replay = append(replay, a)
			if len(replay) > replayCount {
				replay = replay[1:]
			}

			replayMutex.Unlock()

			for _, subscribe := range subscribes {
				subscribe <- item
			}
		}
	}()

	subscribeFunc := func() <-chan Item[A] {
		source := make(chan Item[A])

		ch := make(chan Item[A])

		go func() {
			defer close(ch)

			replayMutex.RLock()
			for _, a := range replay {
				ch <- ItemOf[A](a)
			}
			replayMutex.RUnlock()

			if closed {
				return
			}

			subscribes = append(subscribes, source)

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
