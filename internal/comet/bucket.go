package comet

import (
	"errors"
	"sync"
)

type Bucket struct{
	bct map[string]*Channel

	mu sync.RWMutex
}

func NewBucket(size int) *Bucket {
	return &Bucket{
		bct: make(map[string]*Channel, size),
	}
}

func (b *Bucket) GetChannel(key string) (*Channel, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	ch, ok := b.bct[key]
	if !ok {
		return nil, errors.New("channel not found")
	}
	return ch, nil
}

func (b *Bucket) PutChannel(key string, ch *Channel) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.bct[key] = ch
	return nil
}

func (b *Bucket) DeleteChannel(key string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.bct, key)
	return nil
}