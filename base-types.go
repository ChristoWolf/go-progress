package progress

import (
	"errors"
	"io"
	"time"
)

// Progresser is an interface which defines contracts for interacting with simple progress visualizers.
type Progresser interface {
	// Method for starting progress visualization.
	// Should be called as a goroutine to allow for asynchronous work execution
	// for which its progress is supposed to be visualized.
	Start() error
	// Method for stopping execution of the Start() goroutine via signaling and closing a channel.
	Stop() error
}

// Base type which provides common fields to eassily supply progress visualization logic.
// Should be via struct embedding it as an anonymous field in Progresser implementations.
// Supports arbitrary visualization sinks; everything which implements io.Writer.
type baseProgress struct {
	delay  time.Duration
	sink   io.Writer
	stop   chan struct{}
	ticker *time.Ticker
}

// Method usually used for Progressor implementations (which make use of baseProgress)
// for signaling the progress to stop and closing its signaling channel.
func (p *baseProgress) Stop() error {
	if p.stop == nil {
		return errors.New("no channel to stop")
	}
	select {
	case _, ok := <-p.stop: // Safety check to ensure that the channel is not closed already.
		if !ok {
			return errors.New("channel has already been stopped")
		}
	default:
		p.stop <- struct{}{}
		close(p.stop)
	}
	return nil
}
