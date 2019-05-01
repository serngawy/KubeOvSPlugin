package utils

import (
	"sync"
)

// Broadcaster holds the details of registered listeners
type Broadcaster struct {
	listenerLock sync.RWMutex
	listeners    []Listener
}

// NewBroadcaster returns an instance of Broadcaster object
func NewBroadcaster() *Broadcaster {
	return &Broadcaster{}
}

// Add lets to register a listener
func (b *Broadcaster) Add(listener Listener) {
	b.listenerLock.Lock()
	defer b.listenerLock.Unlock()
	b.listeners = append(b.listeners, listener)
}

// Notify notifies an update to registered listeners
func (b *Broadcaster) Notify(instance interface{}) {
	b.listenerLock.RLock()
	listeners := b.listeners
	b.listenerLock.RUnlock()
	for _, listener := range listeners {
		go listener.OnUpdate(instance)
	}
}
