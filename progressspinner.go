package progress

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// A simple progress spinner for displaying asynchronous activity.
type progressSpinner struct {
	baseProgress
}

// Creates a new instance of a progress spinner.
// Progress visualization can be customized by configuring:
//
// - they delay between visualization frames,
//
// - and the sink where the visualization is written to.
func NewProgressSpinner(delay time.Duration, sink io.Writer) Progresser {
	return &progressSpinner{baseProgress{delay, sink, make(chan struct{}), nil}}
}

// Method for starting progress visualization.
// Should be called as a goroutine to allow for asynchronous work execution
// for which its progress is supposed to be visualized.
func (p *progressSpinner) Start() error {
	if p.ticker != nil {
		return errors.New("spinner has already been started and/or stopped")
	}
	p.ticker = time.NewTicker(p.delay)
	fmt.Fprintf(p.sink, "-")
	defer p.ticker.Stop()
	for {
		for _, spin := range `\|/-` {
			select {
			case <-p.stop:
				fmt.Fprintf(p.sink, "\b")
				return nil
			case <-p.ticker.C:
				fmt.Fprintf(p.sink, "\b%c", spin)
			}
		}
	}
}
