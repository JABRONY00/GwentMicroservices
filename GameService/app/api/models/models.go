package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name    string
	TableID string
	Conn    *websocket.Conn
}

type WaitingClient struct {
	Name string
	Ch   chan struct{}
}

type PlayerPreset struct {
	Race  string
	Stack []uint
}

type ConcMap[T any] struct {
	sync.RWMutex
	Content map[string]T
}

func (c *ConcMap[T]) Init() {
	c.Lock()
	defer c.Unlock()
	c.Content = make(map[string]T)
}

func (c *ConcMap[T]) Set(key string, value T) {
	c.Lock()
	defer c.Unlock()
	c.Content[key] = value
}

func (c *ConcMap[T]) Get(key string) T {
	c.Lock()
	defer c.Unlock()
	value := c.Content[key]
	return value
}

func (c *ConcMap[T]) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.Content, key)
}
