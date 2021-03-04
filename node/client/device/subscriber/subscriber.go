package subscriber

import (
	"context"
	"encoding/base64"
	"github.com/forestvpn/go-api-client/node/models"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Option func(s *Subscriber)

func WithBasicAuth(username, password string) Option {
	return func(s *Subscriber) {
		s.header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	}
}

func WithHeader(key, val string) Option {
	return func(s *Subscriber) {
		s.header.Set(key, val)
	}
}

// Subscriber is a utility for receiving wireguard key events from a message-queue server
type Subscriber struct {
	url    string
	header http.Header
}

func New(url string, options ...Option) *Subscriber {
	s := &Subscriber{url: url, header: http.Header{}}

	for _, option := range options {
		option(s)
	}

	return s
}

const subProtocol = "message-queue-v1"

// Subscribe establishes a websocket connection for a message-queue channel, and emits messages on the given channel
func (s *Subscriber) Subscribe(ctx context.Context, channel chan<- models.DeviceEvent) error {
	err := s.connect(ctx, channel)

	if err != nil {
		return err
	}

	return nil
}

func (s *Subscriber) connect(ctx context.Context, channel chan<- models.DeviceEvent) error {
	conn, _, err := websocket.Dial(ctx, s.url, &websocket.DialOptions{
		Subprotocols: []string{subProtocol},
		HTTPHeader:   s.header,
	})

	if err != nil {
		return err
	}

	go s.read(ctx, channel, conn)

	return nil
}

func (s *Subscriber) read(ctx context.Context, channel chan<- models.DeviceEvent, conn *websocket.Conn) {
	for {
		v := models.DeviceEvent{}
		err := wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println("error reading from websocket, reconnecting", err)

			// Make sure the connection is closed
			conn.Close(websocket.StatusInternalError, "")

			// Start attempting to reconnect
			go s.reconnect(ctx, channel)

			return
		}

		channel <- v
	}
}

func (s *Subscriber) reconnect(ctx context.Context, channel chan<- models.DeviceEvent) {
	// Sleep
	time.Sleep(time.Second)

	// Attempt to create a new connection
	err := s.connect(ctx, channel)
	if err != nil {
		go s.reconnect(ctx, channel)
	} else {
		log.Println("successfully reconnected to websocket")
	}
}
