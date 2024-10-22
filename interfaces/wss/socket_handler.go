package wss

import (
	"github.com/biangacila/telco-websock/infrastructure/websockets"
	"github.com/gorilla/websocket"
	"net/http"
)

// WebSocketHandler handles WebSocket connections
type WebSocketHandler struct {
	wsManager *websockets.WebSocketManager
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(wsManager *websockets.WebSocketManager) *WebSocketHandler {
	return &WebSocketHandler{wsManager: wsManager}
}

// ServeWebSocket handles new WebSocket connections
func (h *WebSocketHandler) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	userCode := r.URL.Query().Get("userCode")
	if userCode == "" {
		http.Error(w, "Missing userCode", http.StatusBadRequest)
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	h.wsManager.AddConnection(userCode, conn)

	// Listen for WebSocket disconnection
	go func() {
		defer conn.Close()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				h.wsManager.RemoveConnection(userCode)
				break
			}
		}
	}()
}
