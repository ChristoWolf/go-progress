package progress_test

import (
	"os"
	"testing"
	"time"

	"github.com/christowolf/go-progress"
)

func TestProgressDotsStart(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		busyTime time.Duration
		message  string
	}{
		{time.Microsecond, 50 * time.Millisecond, "test message"},
	}
	for _, row := range data {
		spy := writerSpy{}
		prog := progress.NewProgressDots(row.delay, &spy, row.message)
		prog.Start()
		time.Sleep(row.busyTime)
		prog.Stop()
		want := row.message
		for i := 0; i < spy.ticks-1; i++ { // One less to take extra call (to print the message text) to io.Writer.Write() into account.
			want += "."
		}
		if got := spy.string; got != want || spy.ticks <= 0 {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

// Simple example which demonstrates how to use the progress dots visualization.
// Change the variables with 'example' prefix to customize the example.
// No output is checked here due to tick randomness.
func ExampleProgressDots() {
	exampleDelay := 100 * time.Millisecond
	exampleBusyTime := 1 * time.Second
	exampleMessage := "working"
	prog := progress.NewProgressDots(exampleDelay, os.Stdout, exampleMessage)
	prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
}
