package monitoring

/*
import (
	"context"
	"log"
	"time"

	"github.com/gagliardetto/solana-go"
)

type Poller interface {
	Start(context.Context)
	Updates() <-chan interface{}
}

func NewPoller(
	account solana.PublicKey,
	reader AccountReader,
	fetchInterval time.Duration,
	bufferCapacity uint32,
) Poller {
	return &solanaPollerImpl{
		account,
		reader,
		fetchInterval,
		bufferCapacity,
		make(chan interface{}, bufferCapacity),
	}
}

type solanaPollerImpl struct {
	account        solana.PublicKey
	reader         AccountReader
	fetchInterval  time.Duration
	bufferCapacity uint32
	updates        chan interface{}
}

// Start should be executed as a goroutine
func (s *solanaPollerImpl) Start(ctx context.Context) {
	// Fetch initial data
	data, err := s.reader.Read(ctx, s.account)
	if err != nil {
		// TODO (dru) handle error, don't return
		log.Printf("failed to read account (address=%s), data: %v", s.account.String(), err)
	}
	s.updates <- data

	for {
		timer := time.NewTimer(s.fetchInterval)
		select {
		case <-timer.C:
			data, err := s.reader.Read(ctx, s.account)
			if err != nil {
				// TODO (dru) handle error
				continue
			}
			s.updates <- data
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return
		}
	}
}

func (s *solanaPollerImpl) Updates() <-chan interface{} {
	return s.updates
}
*/
