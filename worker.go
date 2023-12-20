package rex

func doWithPool(ctx Context, poolSize int, do func()) {
	for i := 0; i < poolSize; i++ {
		go do()
	}
}
