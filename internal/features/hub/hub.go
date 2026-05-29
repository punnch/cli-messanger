package hub

import (
	"net"
	"sync"
)

type BroadcastMsg struct {
	RoomID  string
	Payload []byte
}

type Hub struct {
	mu        sync.RWMutex
	rooms     map[string][]net.Conn
	Broadcast chan BroadcastMsg
}

func NewHub() *Hub {
	return &Hub{
		mu:        sync.RWMutex{},
		rooms:     make(map[string][]net.Conn),
		Broadcast: make(chan BroadcastMsg, 256),
	}
}

func (h *Hub) SendMessage(
	roomID string,
	payload []byte,
) {
	h.Broadcast <- BroadcastMsg{
		RoomID:  roomID,
		Payload: payload,
	}
}

func (h *Hub) Register(
	roomID string,
	conn net.Conn,
) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.rooms[roomID] = append(h.rooms[roomID], conn)
}

func (h *Hub) Unregister(
	roomID string,
	conn net.Conn,
) {
	h.mu.Lock()
	defer h.mu.Unlock()

	conns := h.rooms[roomID]
	for i, c := range conns {
		if c == conn {
			h.rooms[roomID] = append(conns[:i], conns[i+1:]...)
			return
		}
	}
}

// Run() — one goroutine, fans out to all conns in a room.
func (h *Hub) Run() {
	for msg := range h.Broadcast {
		h.mu.RLock()
		for _, conn := range h.rooms[msg.RoomID] {
			_, _ = conn.Write(msg.Payload)
		}
		h.mu.RUnlock()
	}
}
