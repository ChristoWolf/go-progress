package progress_test

// Type for mocking io.Writer in tests.
// Contains a string for writing to and a counter which increments with each call to the Write method.
type writerSpy struct {
	string
	ticks int
}

// Write method implementation to enable io.Writer for this mock.
// Appends the bytes to write to its embedded string and increments the 'ticks' counter.
func (w *writerSpy) Write(p []byte) (n int, err error) {
	w.string += string(p)
	w.ticks++
	return len(p), nil
}
