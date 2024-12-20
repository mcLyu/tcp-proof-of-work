package tcpserver

import "sync/atomic"

type ConnectionsCounter struct {
	n atomic.Int32
}

func (c *ConnectionsCounter) Increment() {
	c.n.Add(1)
}

func (c *ConnectionsCounter) Decrement() {
	c.n.Add(-1)
}

func (c *ConnectionsCounter) Count() int32 {
	return c.n.Load()
}
