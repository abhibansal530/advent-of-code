package utils

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func AssertEvalsToTrue(f func() bool) {
	if !f() {
		panic("got false")
	}
}
