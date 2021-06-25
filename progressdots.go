package progress

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

// A simple progress spinner for displaying asynchronous activity.
type progressDots struct {
	baseProgress
	message string
}

// Creates a new instance of a progress dots visualization struct.
// Progress visualization can be customized by configuring:
//
// - they delay between visualization frames,
//
// - and the sink where the visualization is written to.
func NewProgressDots(delay time.Duration, sink io.Writer, message string) Progresser {
	return &progressDots{baseProgress{delay, sink, make(chan struct{}), nil, sync.WaitGroup{}}, message}
}

// Starts concurrent progress dots visualization.
// Execution of the caller goroutine continues and progress visualization may be stoped using Stop().
func (p *progressDots) Start() error {
	if p.ticker != nil {
		return errors.New("progress dots visualization has already been started and/or stopped")
	}
	p.ticker = time.NewTicker(p.delay)
	p.wg.Add(1)
	go p.work()
	return nil
}

// Internally used method containing actual progress dots visualization logic.
// Should be called as a goroutine during Start() of a Progresser.
func (p *progressDots) work() {
	defer p.wg.Done()
	fmt.Fprint(p.sink, p.message)
	for {
		select {
		case <-p.stop:
			return
		case <-p.ticker.C:
			fmt.Fprint(p.sink, ".")
		}
	}
}
