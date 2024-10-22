package websockets

import (
	"github.com/gorilla/websocket"
	"sync"
)

// WebSocketManager manages WebSocket connections
type WebSocketManager struct {
	connections map[string]*websocket.Conn
	lock        sync.RWMutex
}

// NewWebSocketManager creates a new WebSocket manager
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string]*websocket.Conn),
	}
}

// AddConnection adds a WebSocket connection for a given userCode
func (manager *WebSocketManager) AddConnection(userCode string, conn *websocket.Conn) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.connections[userCode] = conn
}

// RemoveConnection removes a WebSocket connection
func (manager *WebSocketManager) RemoveConnection(userCode string) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	delete(manager.connections, userCode)
}

// GetConnection retrieves the WebSocket connection for a given userCode
func (manager *WebSocketManager) GetConnection(userCode string) (*websocket.Conn, bool) {
	manager.lock.RLock()
	defer manager.lock.RUnlock()
	conn, exists := manager.connections[userCode]
	return conn, exists
}
