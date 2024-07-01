package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name    string
	Conn    *websocket.Conn
	TableID string
}

type WaitingClient struct {
	Name string
	Ch   chan bool
}

type ConcMap struct {
	sync.RWMutex
	Content map[string]interface{}
}

func (c *ConcMap) Init() {
	c.Lock()
	defer c.Unlock()
	c.Content = make(map[string]interface{})
}

func (c *ConcMap) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.Content[key] = value
}

func (c *ConcMap) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()
	value := c.Content[key]
	return value
}

func (c *ConcMap) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.Content, key)
}
