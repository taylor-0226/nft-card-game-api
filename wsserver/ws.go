package wsserver

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Server struct {
	ID            int64
	Name          string
	Ctx           context.Context
	handleMessage func(Message) error
	handleError   func(error)
	mu            sync.RWMutex
	config        ServerConfig
	Clients       map[int64]*Client
	queue         []Message
	// conn          net.Conn
}

type Client struct {
	Id       int64
	UserId   int64
	Platform string
	Conn     net.Conn
}
type ServerConfig struct {
	freq       time.Duration
	maxThreads int
}

type MessageType string
type Message struct {
	MessageType MessageType
	Data        interface{}
}

var Global *Server

func Init() {
	Global = &Server{
		Ctx:     context.Background(),
		Clients: make(map[int64]*Client),
		config: ServerConfig{
			freq:       1 * time.Second,
			maxThreads: 10,
		},
	}

	// Global.Start()
}

func (s *Server) AddClient(user_id int64, conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Clients[user_id] = &Client{
		Id:     user_id,
		Conn:   conn,
		UserId: user_id,
	}
	s.Listen(conn)
}

func (s *Server) Listen(conn net.Conn) {
	for {
		select {
		case <-s.Ctx.Done():
			return
		default:
			b, op, err := wsutil.ReadClientData(conn)
			log.Printf("recv: %d - %s", op, string(b))
			msg := Message{}
			err = json.Unmarshal(b, &msg)
			if err != nil {
				s.handleError(err)
			}
			err = s.handleMessage(msg)
			if err != nil {
				s.handleError(err)
			}
		}
	}
}

func (s *Server) sendMessage(conn net.Conn, msg Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	js, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	wsutil.WriteServerMessage(conn, ws.OpText, js)

	return nil
}

func (s *Server) Broadcast(msg Message) {
	for _, c := range s.Clients {
		s.sendMessage(c.Conn, msg)
	}
}

const (
	MESSAGE_TYPE_AUTHENTICATE = MessageType("AUTHENTICATE")
	MESSAGE_TYPE_PING         = MessageType("PING")
	MESSAGE_TYPE_PONG         = MessageType("PONG")
	MESSAGE_TYPE_ERROR        = MessageType("ERROR")
	MESSAGE_TYPE_EVENT        = MessageType("EVENT")
	// MESSAGE_TYPE_MESSAGE      = MessageType("MESSAGE")
	MESSAGE_TYPE_CLOSE = MessageType("CLOSE")
)
