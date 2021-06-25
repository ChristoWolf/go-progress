package progress

import (
	"errors"
	"io"
	"sync"
	"time"
)

// Progresser is an interface which defines contracts for interacting with simple progress visualizers.
type Progresser interface {
	// Method for starting concurrent progress visualization.
	// Execution of the caller goroutine continues and progress visualization may be stoped using Stop().
	Start() error
	// Method for stopping execution of goroutines triggered by Start() via signaling and closing related channels.
	Stop() error
	// Internally used method containing actual progress visualization logic specific to each Progresser implementation.
	// Should be called as a goroutine during Start() of a Progresser.
	work()
}

// Base type which provides common fields to eassily supply progress visualization logic.
// Should be via struct embedding it as an anonymous field in Progresser implementations.
// Supports arbitrary visualization sinks; everything which implements io.Writer.
type baseProgress struct {
	delay  time.Duration
	sink   io.Writer
	stop   chan struct{}
	ticker *time.Ticker
	wg     sync.WaitGroup
}

// Method usually used for Progressor implementations (which make use of baseProgress)
// for signaling the progress to stop and closing its channels.
func (p *baseProgress) Stop() error {
	if p.stop == nil {
		return errors.New("no channel to stop")
	}
	if p.ticker == nil {
		return errors.New("no ticker to stop")
	}
	select {
	case _, ok := <-p.stop: // Safety check to ensure that the channel is not closed already.
		if !ok {
			return errors.New("channel has already been stopped")
		}
	case _, ok := <-p.ticker.C: // Safety check to ensure that the ticker is not stopped already.
		if !ok {
			return errors.New("ticker has already been stopped")
		}
	default:
		p.ticker.Stop()
		p.stop <- struct{}{}
		close(p.stop)
		p.ticker = nil
	}
	p.wg.Wait()
	return nil
}
