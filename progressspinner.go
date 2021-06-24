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

func (p *progressSpinner) Start() error {
	if p.ticker != nil {
		return errors.New("Progresser has already been started and/or stopped")
	}
	p.ticker = time.NewTicker(p.delay)
	go p.work()
	return nil
}

// Internally used method containing actual progress spinner visualization logic.
// Should be called as a goroutine during Start() of a Progresser.
func (p *progressSpinner) work() {
	fmt.Fprintf(p.sink, "-")
	for {
		for _, spin := range `\|/-` {
			select {
			case <-p.stop:
				fmt.Fprintf(p.sink, "\b")
				return
			case <-p.ticker.C:
				fmt.Fprintf(p.sink, "\b%c", spin)
			}
		}
	}
}
