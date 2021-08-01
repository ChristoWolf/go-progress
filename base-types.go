package progress

import (
	"errors"
	"io"
	"sync"
	"time"
)

// Base type which provides common fields to easily supply progress visualization logic.
// Should be via struct embedding it as an anonymous field in concrete progress visualization struct implementations.
// Supports arbitrary visualization sinks; everything which implements io.Writer.
type baseProgress struct {
	delay  time.Duration
	sink   io.Writer
	stop   chan struct{}
	ticker *time.Ticker
	wg     sync.WaitGroup
}

// Method for signaling the progress visualizer
// (which make use of baseProgress) to stop and close its channels.
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
