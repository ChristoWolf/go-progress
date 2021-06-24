package progress_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/christowolf/go-progress"
)

func TestProgressSpinnerStart(t *testing.T) {
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
		want := spy.ticks - 1                                                // One less to take extra call to io.Writer.Write() into account.
		if got := strings.Count(spy.string, "\b"); got != want || got <= 0 { // TODO: Could be improved by manually constructing 'want' from spy.ticks.
			t.Errorf("got: %d, want: %d", got, want)
		}
	}
}

// Simple example which demonstrates how to use the progress spinner.
// Change the variables with 'example' prefix to customize the example.
// No output is checked here due to tick randomness.
func ExampleProgressSpinner() {
	exampleDelay := 100 * time.Millisecond
	exampleBusyTime := 2 * time.Second
	prog := progress.NewProgressSpinner(exampleDelay, os.Stdout)
	prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
}
