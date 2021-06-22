package progress_test

import (
	"testing"
	"time"

	"github.com/christowolf/go-progress"
)

func TestStart(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		busyTime time.Duration
		want     int
	}{
		{time.Microsecond, time.Millisecond, 0},
	}
	for _, row := range data {
		spy := writerSpy{}
		prog := progress.NewProgressSpinner(row.delay, &spy)
		go prog.Start()
		time.Sleep(row.busyTime)
		prog.Stop()
		if got := spy.ticks; got <= row.want {
			t.Errorf("got: %d, want: > %d", got, row.want)
		}
	}
}

// Simple example which demonstrates how to use the progress spinner.
// Change the variables with 'example' prefix to custimze the example.
func Example() {
	exampleWriter := &writerSpy{}
	exampleDelay := 100 * time.Millisecond
	exampleBusyTime := 10 * time.Second
	prog := progress.NewProgressSpinner(exampleDelay, exampleWriter)
	go prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
	// Output:
	//
}
