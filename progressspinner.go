package progress

import (
	"fmt"
	"io"
	"time"
)

// A simple console spinner for displaying asynchronous activity.
type progressSpinner struct {
	baseProgress
}

func NewProgressSpinner(delay time.Duration, sink io.Writer) Progresser {
	return &progressSpinner{baseProgress{delay, sink, make(chan struct{}), nil}}
}

func (p *progressSpinner) Start() {
	p.ticker = time.NewTicker(p.delay)
	defer p.ticker.Stop()
	for {
		for _, spin := range `-\|/` {
			select {
			case <-p.stop:
				fmt.Fprintf(p.sink, "\r")
				return
			case <-p.ticker.C:
				fmt.Fprintf(p.sink, "\r%c", spin)
			}
		}
	}
}
