package progress_test

import (
	"testing"
	"time"

	"github.com/christowolf/go-progress"
)

func TestStart(t *testing.T) {
	var data = []struct {
		delay    time.Duration
		workTime time.Duration
		want     int
	}{
		{time.Microsecond, time.Millisecond, 0},
	}
	for _, row := range data {
		spy := writerSpy{}
		prog := progress.NewProgressSpinner(row.delay, &spy)
		go prog.Start()
		time.Sleep(row.workTime)
		prog.Stop()
		if got := spy.ticks; got <= row.want {
			t.Errorf("got: %d, want: > %d", got, row.want)
		}
	}
}
