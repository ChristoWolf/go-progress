package progress_test

type writerSpy struct {
	string
	ticks int
}

func (w *writerSpy) Write(p []byte) (n int, err error) {
	w.string += string(p)
	w.ticks++
	return len(p), nil
}
