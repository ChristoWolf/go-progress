package progress_test

import (
	"strings"
	"testing"
	"time"

	"github.com/christowolf/go-progress"
)

func TestStart(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		busyTime time.Duration
	}{
		{time.Microsecond, time.Millisecond},
	}
	for _, row := range data {
		spy := writerSpy{}
		prog := progress.NewProgressSpinner(row.delay, &spy)
		prog.Start()
		time.Sleep(row.busyTime)
		prog.Stop()
		want := spy.ticks - 1
		if got := strings.Count(spy.string, "\b"); got != want && got <= 0 {
			t.Errorf("got: %d, want: %d", got, want)
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
	prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
	// Output:
	//
}
