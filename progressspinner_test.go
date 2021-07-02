package progress_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/christowolf/go-progress"
	"github.com/google/go-cmp/cmp"
)

// Tests common use cases of the progress spinner.
func TestProgressSpinnerNewStartStop(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		busyTime time.Duration
	}{
		{time.Microsecond, time.Millisecond},
		{time.Microsecond, 50 * time.Millisecond},
	}
	for _, row := range data { // TODO: Could be parallelized, see https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721.
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

// Tests if Start() errors out when it was already called before without stopping.
func TestProgressSpinnerStartError(t *testing.T) {
	prog := progress.NewProgressSpinner(time.Millisecond, &writerSpy{})
	if got := prog.Start(); !cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: %v", got, nil)
	}
	if got := prog.Start(); cmp.Equal(got, nil) {
		t.Errorf("got: %v, want: not %v", got, nil)
	}
}

// Tests if Stop() errors out when there is nothing to stop.
func TestProgressSpinnerStopError(t *testing.T) {
	prog := progress.NewProgressSpinner(time.Millisecond, &writerSpy{})
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

// Simple example which demonstrates how to use the progress spinner.
// Change the variables with 'example' prefix to customize the example.
// No output is checked here due to tick randomness
// and the fact that '\b' is non-destructive (similarly to e.g. '\r', '\n'),
// i.e. it only changes the output's cursor position,
// but does NOT overwrite/delete the target position's content.
func ExampleNewProgressSpinner() {
	exampleDelay := 100 * time.Millisecond
	exampleBusyTime := 2 * time.Second
	prog := progress.NewProgressSpinner(exampleDelay, os.Stdout)
	prog.Start()
	time.Sleep(exampleBusyTime)
	prog.Stop()
}
