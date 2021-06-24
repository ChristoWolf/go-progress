package progress

import (
	"errors"
	"fmt"
	"io"
	"sync"
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
	return &progressSpinner{baseProgress{delay, sink, make(chan struct{}), nil, sync.WaitGroup{}}}
}

// Starts concurrent spinner progress visualization.
// Execution of the caller goroutine continues and progress visualization may be stoped using Stop().
func (p *progressSpinner) Start() error {
	if p.ticker != nil {
		return errors.New("progress spinner has already been started and/or stopped")
	}
	p.ticker = time.NewTicker(p.delay)
	p.wg.Add(1)
	go p.work()
	return nil
}

// Internally used method containing actual progress spinner visualization logic.
// Should be called as a goroutine during Start() of a Progresser.
func (p *progressSpinner) work() {
	defer p.wg.Done()
	fmt.Fprintf(p.sink, "-")
	for {
		for _, spin := range `\|/-` {
			select {
			case <-p.stop:
				fmt.Fprint(p.sink, "\b")
				return
			case <-p.ticker.C:
				fmt.Fprintf(p.sink, "\b%c", spin)
			}
		}
	}
}
