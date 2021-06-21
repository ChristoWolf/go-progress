package progress

import (
	"io"
	"time"
)

// Progresser is an interface which defines contracts for interacting with simple progress visualizers.
type Progresser interface {
	// Method for starting progress visualization.
	// Should be called as a goroutine to allow for asynchronous work execution
	// for which its progress is supposed to be visualized.
	Start()
	// Method for stopping execution of the Start() goroutine via signaling and closing a channel.
	Stop()
}

type baseProgress struct {
	delay  time.Duration
	sink   io.Writer
	stop   chan struct{}
	ticker *time.Ticker
}

func (p *baseProgress) Stop() {
	if p.stop == nil {
		return
	}
	if _, ok := <-p.stop; !ok {
		return
	}
	p.stop <- struct{}{}
	close(p.stop)
}
