package progress_test

import (
	"os"
	"testing"
	"time"

	"github.com/christowolf/go-progress"
	"github.com/google/go-cmp/cmp"
)

// Tests common use cases of the progress dots visualization.
func TestProgressDotsNewStartStop(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		busyTime time.Duration
		message  string
	}{
		{time.Microsecond, time.Millisecond, "test message 1"},
		{time.Microsecond, 50 * time.Millisecond, "test message 2"},
	}
	for _, row := range data { // TODO: Could be parallelized, see https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721.
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

// Tests if Start() errors out when it was already called before without stopping.
func TestProgressDotsStartError(t *testing.T) {
	prog := progress.NewProgressDots(time.Millisecond, &writerSpy{}, "TEST")
	if got := prog.Start(); !cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: %v", got, nil)
	}
	if got := prog.Start(); cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: not %v", got, nil)
	}
}

// Tests if Stop() errors out when there is nothing to stop.
func TestProgressDotsStopError(t *testing.T) {
	prog := progress.NewProgressDots(time.Millisecond, &writerSpy{}, "TEST")
	if got := prog.Stop(); cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: %v", got, nil)
	}
	prog.Start()
	if got := prog.Stop(); !cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: not %v", got, nil)
	}
	if got := prog.Stop(); cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: not %v", got, nil)
	}
}

// Simple example which demonstrates how to use the progress dots visualization.
// Change the variables with 'example' prefix to customize the example.
// No output is checked here due to tick randomness.
func ExampleNewProgressDots() {
	exampleDelay := 100 * time.Millisecond
	exampleBusyTime := 1 * time.Second
	exampleMessage := "working"
	prog := progress.NewProgressDots(exampleDelay, os.Stdout, exampleMessage)
	prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
}
