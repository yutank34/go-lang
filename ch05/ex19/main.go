package ex19

func f(v int) (result int) {
	defer func() {
		result = recover().(int)
	}()
	panic(v)
}
