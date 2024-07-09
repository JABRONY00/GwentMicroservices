package models

import (
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

type Claims struct {
	jwt.StandardClaims
}

type ResponseData struct {
	Instr string      `json:"instruction"`
	Data  interface{} `json:"data"`
}

type Connection struct {
	Mut *sync.RWMutex
	*websocket.Conn
}

type Client struct {
	Name    string
	TableID string
	Conn    *Connection
}

func (c *Client) SendJson(resp interface{}) error {
	c.Conn.Mut.Lock()
	defer c.Conn.Mut.Unlock()
	err := c.Conn.WriteJSON(resp)
	return err
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
